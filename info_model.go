package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type info_model struct {
	quest        *quest
	active_model tea.Model
	models       info_models
	styles       styles
}

type info_models struct {
	lore_model   lore_model
	legend_model legend_model
}

func (m info_model) refill_models() info_model {
	switch am := m.active_model.(type) {
	case lore_model:
		m.models.lore_model = am
	case legend_model:
		m.models.legend_model = am
	}
	return m
}

func (m info_model) Init() tea.Cmd {
	return nil
}

func (m info_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "3":
			m = m.refill_models()
			m.active_model = m.models.legend_model
		case "4":
			m = m.refill_models()
			m.active_model = m.models.lore_model
		default:
			m.active_model, cmd = m.active_model.Update(msg)
		}
	}
	return m, cmd
}

func (m info_model) View() string {
	upper := quest_header_render(m.quest, m.styles)
	lower := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.styles.info_left_box.Render(m.quest.Description),
		m.styles.info_right_box.Render(m.active_model.View()),
	)
	return lipgloss.JoinVertical(lipgloss.Left, upper, lower)
}

func quest_header_render(q *quest, s styles) string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.PlaceHorizontal(
			66,
			lipgloss.Center,
			s.quest_title.Render(q.Title),
			s.quest_title_whitespace...),
		s.quest_subtitle.Render(q.Subtitle),
	)
}
