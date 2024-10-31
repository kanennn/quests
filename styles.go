package main

import "github.com/charmbracelet/lipgloss"

type styles struct {
	accent lipgloss.Color
	eson   lipgloss.Color
	fg     lipgloss.Color
	bg     lipgloss.Color

	block_border lipgloss.Border

	body lipgloss.Style

	child_title      lipgloss.Style
	child_subtitle   lipgloss.Style
	child_selected   lipgloss.Style
	child_unselected lipgloss.Style

	info_right_box lipgloss.Style
	info_left_box  lipgloss.Style

	quest_title            lipgloss.Style
	quest_title_whitespace []lipgloss.WhitespaceOption
	quest_subtitle         lipgloss.Style
}

func default_styles() styles {
	s := styles{}

	s.accent = lipgloss.Color("7")
	s.eson = lipgloss.Color("4")

	s.fg = s.eson
	s.bg = lipgloss.Color("0")

	s.block_border = lipgloss.Border{
		Top:         "▄",
		Bottom:      "▀",
		Left:        "█",
		Right:       "█",
		TopLeft:     "▄",
		TopRight:    "▄",
		BottomLeft:  "▀",
		BottomRight: "▀",
	}

	s.child_title = lipgloss.NewStyle().Bold(true)
	s.child_subtitle = lipgloss.NewStyle().Faint(true)
	s.child_unselected = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Foreground(s.fg).
		Padding(0, 1).
		Width(34)
	s.child_selected = lipgloss.NewStyle().
		Border(s.block_border).
		BorderForeground(s.fg).
		Background(s.fg).
		Foreground(s.bg).
		Padding(0, 1). // why is this not inherited??
		Inherit(s.child_unselected)

	s.info_left_box = lipgloss.NewStyle().Width(30).Padding(0, 1, 0, 0)
	s.info_right_box = lipgloss.NewStyle().Width(36).Padding(0, 0, 0, 1)

	s.quest_title = lipgloss.NewStyle().
		Bold(true).
		Foreground(s.eson).
		Padding(0, 1).
		Align(lipgloss.Center)
	s.quest_title_whitespace = []lipgloss.WhitespaceOption{
		lipgloss.WithWhitespaceChars("─"),
		lipgloss.WithWhitespaceForeground(s.eson),
	}
	s.quest_subtitle = lipgloss.NewStyle().Faint(true).Align(lipgloss.Center).Padding(0, 0, 1, 0)

	s.body = lipgloss.NewStyle().
		Padding(1, 2).
		Width(70).
		Height(30).Foreground(s.fg)
	return s
}
