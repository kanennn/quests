package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {   
    loaded bool
    tasks *[]Task
    activeModel tea.Model
    viewer *viewer
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

type loadEntry struct { activeClass class}

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
    viewer.capacity = 6
    viewer.editDesc = textinput.New() //TODO this three lines (following) could probably be moved later
    viewer.editDesc.Prompt = ""
    viewer.editDesc.Placeholder = "Description"
    viewer.editDesc.Cursor.SetMode(cursor.CursorBlink)

    return &model{styles: styles, viewer: viewer, tasks: tasks, activeModel: viewer} //this is where you customize the starting scene/model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd //only needed for a couple cases``
    if !m.loaded {
        switch msg := msg.(type) {
        case tea.WindowSizeMsg:
            m.width = msg.Width
            m.height = msg.Height
            return m, nil
        case initData:
            *m.tasks = msg.tasks
            m.viewer.initTasks()
            m.loaded = true
            return m, nil
        default:
            return m, nil
        }    
    } else { switch msg := msg.(type) {
        case entryReturn:
            *m.tasks = append(*m.tasks, msg.t)
            //CreateTasks(*m.tasks)
            m.activeModel = m.viewer
            m.viewer.newCategoryQuest(&(*m.tasks)[len(*m.tasks)-1])
            m.viewer.sortCompletedTasks()
            return m, nil
        case loadEntry:
            entry := newEntry(msg.activeClass)
            entry.styles = m.styles
            *m.viewer = m.activeModel.(viewer)
            m.activeModel = entry
            return m, nil
        case tea.KeyMsg:
            if msg.String() ==  "ctrl+c" {
                CreateTasks(*m.tasks) //TODO later we will need a version that does efficient runtime write outs in case of crashes
                return m, tea.Quit
            } else {
                m.activeModel, cmd = m.activeModel.Update(msg)
                return m, cmd
            }
            
         default: 
            m.activeModel, cmd = m.activeModel.Update(msg)
            return m, cmd
    }}
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
        lipgloss.Place(0, 30, lipgloss.Left, lipgloss.Top, m.activeModel.View()),
    )
    
}

func tui() {
    p := tea.NewProgram(NewModel(),tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there has been an error: %v", err)
        os.Exit(1)
    }
}