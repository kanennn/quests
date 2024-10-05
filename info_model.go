package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type info_model struct {
	quest *quest
}

func (m info_model) name() string {
	return "info"
}

func (m info_model) Init() tea.Cmd {
	return nil
}

func (m info_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m info_model) View() string {
	return m.quest.Description
}
