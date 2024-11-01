package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type legend_model struct {
	quest *quest
	field textinput.Model
	index int
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
			if m.field.Focused() {
				e := new(entry)
				e.time = time.Now()
				e.tag = "yeet"
				e.text = m.field.Value()
				m.quest.legend = append(m.quest.legend, e)
				m.quest.write_legend()
				m.field.Blur()
				m.field.Reset()
			} else {
				m.field.Focus()
			}
		case "ctrl+u":
			m.field.Reset()
			m.field.Blur()
		case "down":
			if m.index < (len(m.quest.legend)-1) && !m.field.Focused() {
				m.index += 1
			} else {
				m.index = 0
			}
		case "up":
			if m.index > 0 && !m.field.Focused() {
				m.index -= 1
			} else {
				m.index += len(m.quest.legend) - 1
			}
		default:
			m.field, cmd = m.field.Update(msg)
		}
	default:
		m.field, cmd = m.field.Update(msg)
	}
	return m, cmd
}

func (m legend_model) View() string {
	var ks []string
	for i, v := range m.quest.legend {
		k := v.time.Format(layout) + " " + v.tag
		if i == m.index {
			k = lipgloss.NewStyle().Bold(true).Render(k)
		}
		ks = append(ks, k)
	}
	var v string
	if m.field.Focused() {
		v = m.field.View()
	} else if len(m.quest.legend) > 0 {
		v = m.quest.legend[m.index].text
	} else {
		v = ""
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, strings.Join(ks, "\n"), v)
}
