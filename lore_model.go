package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type lore_model struct {
	quest *quest
}

func (m lore_model) name() string {
	return "lore"
}

func (m lore_model) Init() tea.Cmd {
	return nil
}

func (m lore_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m lore_model) View() string {
	return string(m.quest.lore)
}
