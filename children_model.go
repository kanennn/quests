package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type children_model struct {
	quest *quest
}

func (m children_model) Init() tea.Cmd {
	return nil
}

func (m children_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m children_model) View() string {
	var right []string
	for _, v := range m.quest.children {
		right = append(right, v.Name+" "+v.Description+"$$$")
	}
	return strings.Join(right, "\n")
}
