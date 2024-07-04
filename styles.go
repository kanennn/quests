package main

import "github.com/charmbracelet/lipgloss"

type Styles struct {
    AccentColor lipgloss.Color
    NormalText lipgloss.Style

    InputField lipgloss.Style
    InfoBox lipgloss.Style
	
    TaskBox lipgloss.Style
    HighLightedTask lipgloss.Style
    NormalTask lipgloss.Style
	CompletedTask lipgloss.Style
	HighLightedCompletedTask lipgloss.Style

    InfoBoxTitle lipgloss.Style
    InfoBoxBreadcrumbs lipgloss.Style
    InfoBoxDesc lipgloss.Style

    CompletedInfoBoxTitle lipgloss.Style
    CompletedInfoBoxBreadcrumbs lipgloss.Style
    CompletedInfoBoxDesc lipgloss.Style
}

func DefaultStyles() *Styles {
    s := new(Styles)
    s.AccentColor = lipgloss.Color("#0ba68a")
    s.InputField = lipgloss.NewStyle().PaddingTop(1).PaddingBottom(1).PaddingRight(2).PaddingLeft(2).Width(80) //.BorderForeground(s.AccentColor).BorderStyle(lipgloss.NormalBorder()).
	
	s.InfoBox = lipgloss.NewStyle().BorderForeground(s.AccentColor).PaddingTop(1).PaddingBottom(1).PaddingRight(2).PaddingLeft(2).Width(48).Height(12)
	s.TaskBox = lipgloss.NewStyle().BorderForeground(s.AccentColor).PaddingTop(1).PaddingBottom(1).PaddingRight(2).PaddingLeft(2).Width(40).Height(12)

    s.NormalTask = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).BorderForeground().BorderStyle(lipgloss.NormalBorder()).Width(34)
    s.HighLightedTask = s.NormalTask.Bold(true).Foreground(s.AccentColor).BorderForeground(s.AccentColor)
    s.CompletedTask = s.NormalTask.Faint(true).Strikethrough(false).BorderForeground(lipgloss.Color("#808080"))
    s.HighLightedCompletedTask = s.CompletedTask.Bold(true).BorderForeground(s.AccentColor)

    s.InfoBoxTitle = lipgloss.NewStyle().Bold(true).Foreground(s.AccentColor).BorderBottom(true).BorderStyle(lipgloss.NormalBorder()).Width(40).PaddingTop(1)
    s.InfoBoxBreadcrumbs = lipgloss.NewStyle().Faint(true).PaddingBottom(1)
    s.InfoBoxDesc = lipgloss.NewStyle().Italic(true).Width(40)

    s.CompletedInfoBoxTitle = s.InfoBoxTitle.UnsetForeground().Faint(true).Strikethrough(true)
    s.CompletedInfoBoxBreadcrumbs = s.InfoBoxBreadcrumbs.UnsetForeground().Faint(true)
    s.CompletedInfoBoxDesc = s.InfoBoxDesc.UnsetForeground().Faint(true)
	


    return s

}