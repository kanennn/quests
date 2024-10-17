package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type children_model struct {
	quest  *quest
	models *models
	index  int
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
	return &m, nil
}

func (m children_model) View() string {
	var right []string
	for i, v := range m.quest.children {
		entry := v.Name + " " + v.Description + "$$$"
		if m.index == i {
			entry = lipgloss.NewStyle().Bold(true).Render(entry)
		}
		right = append(right, entry)
	}
	return strings.Join(right, "\n")
}
