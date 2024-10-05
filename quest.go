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

func (q *quest) open(path string) {
	q.peek(path)
	q.read_legend()
	q.read_lore()
	q.read_children()
}

func (q *quest) read_legend() error {

	file, err := os.Open(filepath.Join(q.dir, "legend.log"))
	defer func() { Check(file.Close()) }()
	if os.IsNotExist(err) {
		return err
	} else {
		Check(err)
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

func (q *quest) read_metadata() error {
	data, err := os.ReadFile(filepath.Join(q.dir, "quest.yml"))
	if os.IsNotExist(err) {
		return err
	} else {
		Check(err)
	}

	// data, err := io.ReadAll(file)
	Check(err)
	err = yaml.Unmarshal(data, q)
	Check(err)

	return err
}

func (q *quest) read_lore() error {
	file, err := os.Open(filepath.Join(q.dir, "lore.md"))
	defer func() { Check(file.Close()) }()
	if os.IsNotExist(err) {
		return err
	} else {
		Check(err)
	}

	data, err := io.ReadAll(file)
	Check(err)
	q.lore = data

	return err
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
			q := new(quest)
			err = q.peek(filepath.Join(q.dir, dir.Name()))
			if err == nil {
				quests = append(quests, q)
			} else if !os.IsNotExist(err) {
				panic(err)
			}
		}
	}
	q.children = quests
}

// func (q *quest) write(E entry) {
// 	bufWriter.Write([]byte(E.time.Format(layout) + "" + E.tag + E.text))
// }
