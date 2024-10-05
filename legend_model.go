package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type legend_model struct {
	quest *quest
}

func (m legend_model) name() string {
	return "legend"
}

func (m legend_model) Init() tea.Cmd {
	return nil
}

func (m legend_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m legend_model) View() string {
	var s []string
	for _, v := range m.quest.legend {
		s = append(s, v.time.Format(layout)+" "+v.tag+" "+v.text)
	}
	return strings.Join(s, "\n")
}
