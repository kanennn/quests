package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type quest struct {
	dir      string
	legend   []*entry
	children []*quest
	parent   *quest
	lore     []byte
	Title    string
	Subtitle string
}

const layout = time.DateTime

type entry struct {
	time time.Time
	text string
	tag  string
}

func (q *quest) get_bit() []byte {
	return bytes.Split(q.lore, []byte("\n"))[0]
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

func open(base string, name string) (*os.File, error) {
	file, err := os.Open(filepath.Join(base, name))
	if os.IsNotExist(err) {
		return os.Open(filepath.Join(base, ".quests", name))
	} else {
		return file, err
	}
}

func read(base string, name string) ([]byte, error) {
	file, err := os.ReadFile(filepath.Join(base, name))
	if os.IsNotExist(err) {
		return os.ReadFile(filepath.Join(base, ".quests", name))
	} else {
		return file, err
	}
}

func (q *quest) read_legend() error {
	file, err := open(q.dir, "legend.log")

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
		bufWriter.Write([]byte(E.time.Format(layout) + " " + E.tag + " " + E.text + "\n"))
	}
	bufWriter.Flush()
}

func (q *quest) read_metadata() error {
	data, err := read(q.dir, "quest.yml")
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
	file, err := open(q.dir, "lore.md")
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
	// var dir string
	// switch q.dir {
	// case "":
	// 	dir = "."
	// default:
	//	dir = q.dir
	// }

	dirs, err := os.ReadDir(q.dir)
	Check(err)
	more_dirs, err := os.ReadDir(filepath.Join(q.dir, ".quests"))

	quests := []*quest{}
	for _, dir := range dirs {
		if dir.IsDir() && dir.Name() != ".quests" {
			n := new(quest)
			err = n.peek(filepath.Join(q.dir, dir.Name()))
			if err == nil {
				quests = append(quests, n)
			} else if !os.IsNotExist(err) {
				panic(err)
			}
		}
	}
	for _, dir := range more_dirs {
		if dir.IsDir() && dir.Name() != ".quests" {
			n := new(quest)
			err = n.peek(filepath.Join(q.dir, ".quests", dir.Name()))
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
	if filepath.Base(dir) == ".quests" {
		dir = filepath.Dir(dir)
	}
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
