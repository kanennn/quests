package main

import (
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewer struct {
    //models *subModels
    tasks *[]Task
    categories map[class][]*Task
    enabledTasks []*Task
    viewClass class
    styles *Styles
    index int
    view int
    capacity int
}

func (m viewer) Init() tea.Cmd {
    return nil
}

func (m *viewer) updateEnabledTasks() {
    var a []*Task
    for i, t := range *m.tasks {
        if !t.Hidden {a = append(a, &(*m.tasks)[i])}
    }
    m.enabledTasks = a
}

// This can be called when a new item is created
func (m *viewer) newCategoryQuest(t *Task) {
    m.categories[t.Class] = append(m.categories[t.Class], t)
        }


func (m *viewer) updateTaskCategories() {
    m.categories = make(map[class][]*Task)
    for _, t := range m.enabledTasks {
        m.newCategoryQuest(t) 
    }
}

func (m *viewer) sortCompletedTasks() {
    var i class
    for i = 0; i < classLen; i++ {
        slices.SortStableFunc(m.categories[i],  func(a, b *Task) int {
            switch {
            case a.Completed == b.Completed: return 0
            case a.Completed, !b.Completed: return 1
            case !a.Completed, b.Completed: return -1
            default: panic("sortCompletedTasks has encountered an impossible scenario")
            }
        })
    }
}

// This should be called anytime tasks are imported from file, or when an item is hidden
func (m *viewer) initTasks() { 
    m.updateTasks()
    //m.sortCompletedTasks()
}

func (m *viewer) updateTasks() {
    m.updateEnabledTasks()
    m.updateTaskCategories()
    m.sortCompletedTasks() //TODO i need this for some weird reason. i didn't think i did probably bc removing or adding a task never requires sorting. however updating categories messing with all the order of tasks
}

func (m viewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd 
    switch msg := msg.(type) {

        case tea.KeyMsg:    
            switch msg.String() {
                
            case "up":
                if m.index > 0 {
                    m.index--
                }
                if m.index == m.view+1 && m.view > 0 {
                    m.view--
                }
            case "down":
                if m.index < len(m.categories[m.viewClass])-1 {
                    m.index++
                }
                if m.index == m.view+m.capacity-1 && m.view+m.capacity < len(m.categories[m.viewClass]) {
                    m.view++
                }
            case "c":
                if len(m.categories[m.viewClass]) > 0 {
                    (*(m.categories[m.viewClass][m.index])).Completed = !(*(m.categories[m.viewClass])[m.index]).Completed
                    m.sortCompletedTasks()
                }
            case "x":
                if len(m.categories[m.viewClass]) > 0 {
                    (*(m.categories[m.viewClass])[m.index]).Hidden = true
                    m.updateTasks()
                    if m.index+1 > len(m.categories[m.viewClass]) {m.index--}
                    if m.view+m.capacity > len(m.categories[m.viewClass]) && m.view > 0 {m.view--}
                }
                 //this is just initTasks should i change it back to that lol
            case "tab":
                if m.viewClass < classLen-1 {m.viewClass++} else {m.viewClass = 0}
                m.index = 0
                m.view = 0
            case "n":
                cmd = func () tea.Msg {return loadEntry{activeClass: m.viewClass}}
            }   
    }
    return m, cmd
}

func (m viewer) View() string {

    var infoBox string
    var taskList string

    var activeCapacity int

    // there may be a more dynamic way to handle all the fields of the task struct but i'm not sure yet
    // TODO at some point we could use the list bubble element
    if m.capacity > len(m.categories[m.viewClass]) {
        activeCapacity = len(m.categories[m.viewClass])
    } else {
        activeCapacity = m.capacity
    }
    
    // var incompleteTasks []*Task
    // var completeTasks []*Task
    // for _, v := range m.categories[m.viewClass] {
    //     if !v.Completed {incompleteTasks = append(incompleteTasks, v)} else {completeTasks = append(completeTasks, v)}
    // }

    // viewTasks := slices.Concat(incompleteTasks, completeTasks)

    viewTasks := m.categories[m.viewClass]

    for i, task := range viewTasks[m.view:m.view+activeCapacity] {
        
        taskStr := task.Name
        if m.index == i+m.view {
            if !task.Completed {
                taskStr = m.styles.HighLightedTask.Render(taskStr)
            } else {
                taskStr = m.styles.HighLightedCompletedTask.Render(taskStr)
            }
        } else {
            if !task.Completed {
                taskStr = m.styles.NormalTask.Render(taskStr)
            } else {
                taskStr = m.styles.CompletedTask.Render(taskStr)
            }
        }
        taskStr += "\n"
        taskList += taskStr
    }

    if len(viewTasks) > 0 {
        selectedTask :=  (viewTasks)[m.index] //throwing error

        var infoTitle string
        var infoBreadCrumbs string
        var infoBoxDesc string
        if !selectedTask.Completed {
            infoTitle = m.styles.InfoBoxTitle.Render(selectedTask.Name)
            infoBreadCrumbs = m.styles.InfoBoxBreadcrumbs.Render(fmt.Sprintf("%s * %s", selectedTask.POI, selectedTask.Context))
            infoBoxDesc = m.styles.InfoBoxDesc.Render(selectedTask.Description)
        } else {
            infoTitle = m.styles.CompletedInfoBoxTitle.Render(selectedTask.Name)
            infoBreadCrumbs = m.styles.CompletedInfoBoxBreadcrumbs.Render(fmt.Sprintf("%s * %s", selectedTask.POI, selectedTask.Context))
            infoBoxDesc = m.styles.CompletedInfoBoxDesc.Render(selectedTask.Description)
        }
        
        infoBox = lipgloss.JoinVertical(lipgloss.Left, infoTitle, infoBreadCrumbs, infoBoxDesc)
    }
    mainArea := lipgloss.JoinHorizontal(lipgloss.Top ,m.styles.TaskBox.Render(taskList), m.styles.InfoBox.Render(infoBox))
    var classHeader string

    switch m.viewClass {
    case main_quest: classHeader = "Main Quests"
    case side_quest: classHeader = "Side Quests"
    case mini_quest: classHeader = "Mini Quests"
    case sleeping_quest: classHeader = "Sleeping Quests"
    default: panic("Unrecognised class type on task:" + string(m.viewClass))
        
    }
    return lipgloss.JoinVertical(lipgloss.Left, classHeader, mainArea)
}