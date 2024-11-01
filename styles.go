package main

import "github.com/charmbracelet/lipgloss"

type styles struct {
	accent lipgloss.Color
	eson   lipgloss.Color
	fg     lipgloss.Color
	bg     lipgloss.Color

	inside_width int

	block_border lipgloss.Border

	body lipgloss.Style

	child_title      lipgloss.Style
	child_subtitle   lipgloss.Style
	child_selected   lipgloss.Style
	child_unselected lipgloss.Style

	child_info_title    lipgloss.Style
	child_info_subtitle lipgloss.Style
	child_info_box      lipgloss.Style

	quest_title            lipgloss.Style
	quest_title_whitespace []lipgloss.WhitespaceOption
	quest_subtitle         lipgloss.Style

	footer lipgloss.Style
}

func default_styles() styles {
	s := styles{}

	s.accent = lipgloss.Color("7")
	s.eson = lipgloss.Color("4")

	s.fg = s.eson
	s.bg = lipgloss.Color("0")

	s.inside_width = 46

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
		Width(23)
	s.child_selected = lipgloss.NewStyle().
		Border(s.block_border).
		BorderForeground(s.fg).
		Background(s.fg).
		Foreground(s.bg).
		Padding(0, 1).Width(23)

	s.quest_title = lipgloss.NewStyle().
		Bold(true).
		Foreground(s.eson).
		Padding(0, 1)
	s.quest_title_whitespace = []lipgloss.WhitespaceOption{
		lipgloss.WithWhitespaceChars("─"),
		lipgloss.WithWhitespaceForeground(s.eson),
	}
	s.quest_subtitle = lipgloss.NewStyle().Faint(true).Padding(0, 0, 1, 0)

	s.child_info_title = lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.NormalBorder(), false, false, true, false).Width(19)
	s.child_info_subtitle = s.quest_subtitle.Faint(true).Width(19)
	s.child_info_box = lipgloss.NewStyle().Width(23).Padding(0, 2)

	s.footer = lipgloss.NewStyle().
		Width(s.inside_width).
		Align(lipgloss.Center).
		Faint(true).
		Padding(2, 0, 0, 0)

	s.body = lipgloss.NewStyle().
		Padding(1, 2).
		Width(50).
		Height(24).Foreground(s.fg)
	return s
}
