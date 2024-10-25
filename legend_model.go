package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type legend_model struct {
	quest *quest
	field textinput.Model
}

func (m legend_model) Init() tea.Cmd {
	return nil
}

func (m legend_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case textinput.Model:
		m.field = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+e":
			m.field.Focus()
		case "enter":
			e := new(entry)
			e.time = time.Now()
			e.tag = "yeet"
			e.text = m.field.Value()
			m.quest.legend = append(m.quest.legend, e)
			m.quest.write_legend()
			m.field.Blur()
			m.field.Reset()
		case "ctrl+u":
			m.field.Reset()
			m.field.Blur()
		default:
			m.field, cmd = m.field.Update(msg)
		}
	default:
		m.field, cmd = m.field.Update(msg)
	}
	return &m, cmd
}

func (m legend_model) View() string {
	var s []string
	for _, v := range m.quest.legend {
		s = append(s, v.time.Format(layout)+" "+v.tag+" "+v.text)
	}
	if m.field.Focused() {
		s = append(s, m.field.View())
	}
	return strings.Join(s, "\n")
}
