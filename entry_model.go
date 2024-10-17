package main

import (
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type entry_model struct {
	old_quest quest
	new_quest quest
	fields    []field_struct
	index     int
	field     field
	ready     bool
}

// type field interface {
// 	Value() string
// }

// func (m entry_model) switch_field(i interface{}) field {
// 	var f field
// 	switch i.(type) {
// 	case string:
// 		f = m.models.text
// 	}
// 	return f
// }

func (m *entry_model) switch_field() {
	// TODO could make use of field.Reset() for optimization and non-duplication
	switch (m.fields[m.index].typ).(type) {
	case string:
		m.field = new_text_field()

	default:
		panic("No field model available to map to invalid type of field")
	}
}

type field_struct struct {
	name string
	typ  interface{}
	fill func(q *quest, i interface{})
}

func new_entry_model(q *quest) tea.Model {
	m := new(entry_model)
	var str string
	m.old_quest = *q
	m.new_quest = quest{}
	m.fields = []field_struct{
		{name: "name", typ: str, fill: func(q *quest, i interface{}) { q.Name = i.(string) }},
		{name: "description", typ: str, fill: func(q *quest, i interface{}) { q.Description = i.(string) }},
	}
	m.switch_field()

	return m
}

func (m entry_model) Init() tea.Cmd {
	return nil // no need for async init right now
}
func (m entry_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "enter":
			v := m.field.Value()
			m.fields[m.index].fill(&m.new_quest, v)

			if (m.index + 1) >= len(m.fields) {
				m.new_quest.parent = &m.old_quest
				m.new_quest.dir = filepath.Join(m.old_quest.dir, strings.ToLower(strings.ReplaceAll(m.new_quest.Name, " ", "_")))
				m.new_quest.write_all()
				cmd = func() tea.Msg {
					return m.new_quest
				}
			} else {
				m.index += 1
				m.switch_field()
			}

			// BIG tHINK

			// what if we just have separate models for different fields
			// we can spinup each of these for the respective fields and types
			// so just test the type of the field and spinup an entry field, and then pass it the field to fill
			// and these can be embedded side by side
			// and these could be used outside the new-quest function
		default:
			m.field, cmd = m.field.Update(msg)
		}
	default:
		m.field, cmd = m.field.Update(msg)
	}
	return &m, cmd
}

func (m entry_model) View() string { return m.fields[m.index].name + "\n" + m.field.View() }
