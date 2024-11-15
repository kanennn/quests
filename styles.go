package main

import "github.com/charmbracelet/lipgloss"

type styles struct {
	accent lipgloss.Color
	eson   lipgloss.Color
	fg     lipgloss.Color
	bg     lipgloss.Color

	inside_width   int
	inside_height  int
	content_height int

	block_border lipgloss.Border

	body lipgloss.Style

	menu_selected   lipgloss.Style
	menu_unselected lipgloss.Style
	menu_box        lipgloss.Style

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

	main lipgloss.Style
}

func default_styles() styles {
	s := styles{}

	s.accent = lipgloss.Color("7")
	s.eson = lipgloss.Color("4")

	s.fg = s.eson
	s.bg = lipgloss.Color("0")

	s.inside_width = 76
	s.inside_height = 42

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

	s.menu_selected = lipgloss.NewStyle().Padding(0, 3).Bold(true).Foreground(s.fg)
	s.menu_unselected = lipgloss.NewStyle().Padding(0, 3).Faint(true).Foreground(s.fg)
	s.menu_box = lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Width(s.inside_width).
		PaddingBottom(1)

	s.child_title = lipgloss.NewStyle().Bold(true)
	s.child_subtitle = lipgloss.NewStyle().Faint(true)
	s.child_unselected = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(s.fg).
		Foreground(s.fg).
		Padding(0, 1).
		Width(s.inside_width / 2)
	s.child_selected = lipgloss.NewStyle().
		Border(s.block_border).
		BorderForeground(s.fg).
		Background(s.fg).
		Foreground(s.bg).
		Padding(0, 1).Width(s.inside_width / 2)

	s.quest_title = lipgloss.NewStyle().
		Bold(true).
		Foreground(s.fg).
		Padding(0, 1)
	s.quest_title_whitespace = []lipgloss.WhitespaceOption{
		lipgloss.WithWhitespaceChars("─"),
		lipgloss.WithWhitespaceForeground(s.fg),
	}
	s.quest_subtitle = lipgloss.NewStyle().Faint(true).Padding(0, 0, 1, 0).Foreground(s.fg)

	s.child_info_title = lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(s.fg).Width(s.inside_width/2 - 4).Foreground(s.fg)
	s.child_info_subtitle = s.quest_subtitle.Faint(true).
		Width(s.inside_width/2 - 4).
		Foreground(s.fg)
	s.child_info_box = lipgloss.NewStyle().Width(s.inside_width/2 - 2).PaddingLeft(2)

	s.footer = lipgloss.NewStyle().
		Width(s.inside_width).
		Align(lipgloss.Center).
		Faint(true).
		Padding(2, 0, 0, 0).
		Foreground(s.fg) // does everything need to be blue?

	s.content_height = s.inside_height - 3 - 2 - 3

	s.body = lipgloss.NewStyle().
		Padding(1, 2).
		Width(s.inside_width + 4).
		Height(s.inside_height + 2)
	return s
}
