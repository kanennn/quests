package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type lore_model struct {
	quest *quest
	field textinput.Model
}

func (m lore_model) Init() tea.Cmd {
	return nil
}

func (m lore_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case textinput.Model:
		m.field = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+e":
			if m.field.Focused() {
				m.quest.lore = []byte(m.field.Value())
				m.quest.write_lore()
				m.field.Blur()
			} else {
				m.field.SetValue(string(m.quest.lore))
				m.field.Focus()
			}
		case "ctrl+u":
			m.field.Reset()
			m.field.Blur()
		default:
			m.field, cmd = m.field.Update(msg)
		}
	default:
		m.field, cmd = m.field.Update(msg)
	}
	return m, cmd
}

func (m lore_model) View() string {
	if m.field.Focused() {
		return m.field.View()
	} else {
		return string(m.quest.lore)
	}
}
