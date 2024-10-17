package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type quest struct {
	dir         string
	legend      []*entry
	children    []*quest
	parent      *quest
	lore        []byte
	Name        string
	Description string
}

const layout = time.DateTime

type entry struct {
	time time.Time
	text string
	tag  string
}

func (q *quest) peek(path string) error {
	q.dir = path

	err := q.read_metadata()
	return err
}

func (q *quest) open() {
	q.read_legend()
	q.read_lore()
	q.read_children()
	q.read_parent()
}

func (q *quest) read_legend() error {

	file, err := os.Open(filepath.Join(q.dir, "legend.log"))

	if os.IsNotExist(err) {
		return err
	} else {
		Check(err)
		defer func() { err := file.Close(); Check(err) }()
	}

	bufScanner := bufio.NewScanner(file)
	var entries []*entry

	for bufScanner.Scan() {
		e := new(entry)
		s := strings.Split(bufScanner.Text(), " ")

		var err error
		e.time, err = time.Parse(layout, strings.Join(s[:2], " "))
		Check(err)

		e.tag = s[2]
		e.text = strings.Join(s[3:], " ")
		entries = append(entries, e)
	}
	Check(bufScanner.Err())

	q.legend = entries
	return err
}

func (q *quest) write_legend() {

	file, err := os.Create(filepath.Join(q.dir, "legend.log"))
	Check(err)

	defer func() { err := file.Close(); Check(err) }()

	bufWriter := bufio.NewWriter(file)

	for _, E := range q.legend {
		bufWriter.Write([]byte(E.time.Format(layout) + " " + E.tag + " " + E.text))
	}

}

func (q *quest) read_metadata() error {
	data, err := os.ReadFile(filepath.Join(q.dir, "quest.yml"))
	if os.IsNotExist(err) {
		return err
	} else {
		Check(err)
	}

	Check(err)
	err = yaml.Unmarshal(data, q)
	Check(err)

	return err
}

func (q *quest) write_metadata() {
	file, err := os.Create(filepath.Join(q.dir, "quest.yml"))
	Check(err)

	defer func() { err := file.Close(); Check(err) }()

	data, err := yaml.Marshal(q)
	Check(err)

	file.Write(data)
}

func (q *quest) read_lore() error {
	file, err := os.Open(filepath.Join(q.dir, "lore.md"))
	if os.IsNotExist(err) {
		return err
	} else {
		Check(err)
		defer func() { err := file.Close(); Check(err) }()
	}

	data, err := io.ReadAll(file)
	Check(err)
	q.lore = data

	return err
}

func (q *quest) write_lore() {
	file, err := os.Create(filepath.Join(q.dir, "lore.md"))
	Check(err)
	defer func() { err := file.Close(); Check(err) }()

	_, err = file.Write(q.lore)
	Check(err)
}

func (q *quest) read_children() {
	var dir string
	switch q.dir {
	case "":
		dir = "."
	default:
		dir = q.dir
	}
	dirs, err := os.ReadDir(dir)
	Check(err)

	quests := []*quest{}
	for _, dir := range dirs {
		if dir.IsDir() {
			n := new(quest)
			err = n.peek(filepath.Join(q.dir, dir.Name()))
			if err == nil {
				quests = append(quests, n)
			} else if !os.IsNotExist(err) {
				panic(err)
			}
		}
	}
	q.children = quests
}

func (q *quest) read_parent() {
	dir := filepath.Dir(q.dir)

	p := new(quest)
	err := p.peek(dir)

	if !os.IsNotExist(err) {
		Check(err)
		q.parent = p
	}
}

func (q *quest) write_dir() {
	err := os.Mkdir(q.dir, os.ModePerm)
	Check(err)
}

func (q *quest) write_all() {
	q.write_dir()
	q.write_metadata()
	q.write_lore()
	q.write_legend()
}
