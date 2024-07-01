package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// type subModels struct {
//     viewer tea.Model
//     entry tea.Model
// }

type model struct {   
    loaded bool
    tasks *[]Task
    activeModel tea.Model
    //subModels *subModels
    viewer tea.Model
    width int
    height int     
    styles *Styles
}

type initData struct {
    tasks []Task
}

type entryReturn struct {
    t Task
}

type loadEntry struct {}

type Styles struct {
    BorderColor lipgloss.Color
    InputField lipgloss.Style
    InfoBox lipgloss.Style
    TaskBox lipgloss.Style
    HighLightedTask lipgloss.Style
    NormalTask lipgloss.Style
}

func DefaultStyles() *Styles {
    s := new(Styles)
    s.BorderColor = lipgloss.Color("#ee40f1")
    s.InputField = lipgloss.NewStyle().PaddingTop(1).PaddingBottom(1).PaddingRight(2).PaddingLeft(2).Width(80) //.BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).
    s.InfoBox = lipgloss.NewStyle().BorderForeground(s.BorderColor).PaddingTop(1).PaddingBottom(1).PaddingRight(2).PaddingLeft(2).Width(48).Height(12)
    s.TaskBox = lipgloss.NewStyle().BorderForeground(s.BorderColor).PaddingTop(1).PaddingBottom(1).PaddingRight(2).PaddingLeft(2).Width(30).Height(12)
    s.HighLightedTask = lipgloss.NewStyle().Bold(true).Foreground(s.BorderColor)
    return s

}

func getInitData() tea.Msg {
	return initData{tasks: getTasks()}
}


func (m model) Init() (tea.Cmd) {
	return getInitData
}

func NewModel() *model {
    styles := DefaultStyles()
    tasks := new([]Task)

    viewer := new(viewer)
    viewer.styles = styles
    viewer.tasks = tasks

    return &model{styles: styles, viewer: viewer, tasks: tasks, activeModel: viewer} //this is where you customize the starting scene/model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
        
        case tea.WindowSizeMsg:
            m.width = msg.Width
            m.height = msg.Height
            return m, nil
        case initData:
            *m.tasks = msg.tasks
            m.loaded = true
            return m, nil
        case entryReturn:
            *m.tasks = append(*m.tasks, msg.t)
            m.activeModel = m.viewer
            return m, nil
        case loadEntry:
            entry := newEntry()
            entry.styles = m.styles
            m.activeModel = entry
            return m, nil
        case tea.KeyMsg:
            if msg.String() ==  "ctrl+c" {
                return m, tea.Quit
            } else {
                m.activeModel, cmd = m.activeModel.Update(msg)
                return m, cmd
            }
            
         default: 
            m.activeModel, cmd = m.activeModel.Update(msg)
            return m, cmd
    }
}

func (m model) View() string {
    m.activeModel.View()
    if m.width == 0 {
        return "loading window"
    }

    if !m.loaded {
        return "loading application"
    }
    return lipgloss.Place(
        m.width,
        m.height,
        lipgloss.Center,
        lipgloss.Center,
        m.activeModel.View(), 
    )
    
}

func tui() {
    p := tea.NewProgram(NewModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there has been an error: %v", err)
        os.Exit(1)
    }
}