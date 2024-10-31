package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type children_model struct {
	quest *quest
	index int
	s     styles
}

func (m children_model) Init() tea.Cmd {
	return nil
}

func (m children_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "down":
			if m.index < (len((*m.quest).children) - 1) {
				m.index += 1
			} else {
				m.index = 0
			}
		case "up":
			if m.index > 0 {
				m.index -= 1
			} else {
				m.index += len(m.quest.children) - 1
			}
		case "enter":
			if len(m.quest.children) > 0 {
				return &m, func() tea.Msg {
					c := m.quest.children[m.index]
					c.open()
					return *c
				}
			}
		}
	}
	return m, nil
}

func (m children_model) View() string {
	var entries []string
	var lower string
	for i, v := range m.quest.children {
		var entry string
		if m.index == i {
			entry = m.s.child_selected.Render(
				m.s.child_title.Render(v.Title) + "\n" + v.Subtitle,
			)
		} else {
			entry = m.s.child_unselected.Render(m.s.child_title.Render(v.Title) + "\n" + m.s.child_subtitle.Render(v.Subtitle))
		}
		entries = append(entries, entry)
	}
	if len(m.quest.children) > 0 {
		lower = lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Left, entries...),
			m.quest.children[m.index].Description,
		)
	} else {
		lower = ""
	}
	upper := quest_header_render(m.quest, m.s)
	return lipgloss.JoinVertical(lipgloss.Left, upper, lower)
}
