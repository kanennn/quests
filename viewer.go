package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewer struct {
    //models *subModels
    tasks *[]Task
    styles *Styles
    index int
    view int
    capacity int
}

func (m viewer) Init() tea.Cmd {
    return nil
}
    
func (m viewer) activeTasks() [](*Task) {
    var a []*Task
    for i, t := range *m.tasks {
        if !t.Hidden {a = append(a, &(*m.tasks)[i])}
    }
    
    return a
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
                if m.index < len(m.activeTasks())-1 {
                    m.index++
                }
                if m.index == m.view+m.capacity-1 && m.view+m.capacity < len(m.activeTasks()) {
                    m.view++
                }
            case "c":
                (*(m.activeTasks())[m.index]).Completed = !(*(m.activeTasks())[m.index]).Completed
                sortTasks(m.tasks)
            case "x":
                (*(m.activeTasks())[m.index]).Hidden = true
            case "n":
                return m, func () tea.Msg {return loadEntry{}}
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
    if m.capacity > len(m.activeTasks()) {
        activeCapacity = len(m.activeTasks())
    } else {
        activeCapacity = m.capacity
    }
    for i, task := range (m.activeTasks())[m.view:m.view+activeCapacity] {

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

    if len(m.activeTasks()) > 0 {
        selectedTask :=  (m.activeTasks())[m.index]

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

    return lipgloss.JoinHorizontal(lipgloss.Top, m.styles.TaskBox.Render(taskList), m.styles.InfoBox.Render(infoBox))
}
// var infoBox string
// var taskList string

// for i, task := range *m.tasks {

//     taskStr := task.Name
//     if m.index == i {
//         taskStr = m.styles.HighLightedTask.Render(taskStr)
//     }
//     taskStr += "\n"
//     taskList += taskStr
// }

// if len(*m.tasks) > 0 {
//     selectedTask :=  (*m.tasks)[m.index]
//     infoBox = fmt.Sprintf("[*] %s\n[@] %s\n\n%s", selectedTask.Context, selectedTask.POI, selectedTask.Description)
// }
// return lipgloss.Place(
//     m.mother.width,
//     m.mother.height,
//     lipgloss.Center,
//     lipgloss.Center,
//     lipgloss.JoinVertical(
//         lipgloss.Left, 
//         lipgloss.JoinHorizontal(
//             lipgloss.Top, 
//             m.styles.TaskBox.Render(taskList), 
//         m.styles.InputField.Render(m.input.View()),
//     ),
// )