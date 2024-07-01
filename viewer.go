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
}

func (m viewer) Init() tea.Cmd {
    return nil
}



func (m viewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

        case tea.KeyMsg:    
            switch msg.String() {
                
            case "up":
                if m.index > 0 {
                    m.index--
                }
            case "down":
                if m.index < len(*m.tasks)-1 {
                    m.index++
                }
            case "n":
                return m, func () tea.Msg {return loadEntry{}}
            }   
    }
    return m, nil
}

func (m viewer) View() string {

    var infoBox string
    var taskList string

    // TODO at some point we could use the list bubble element
    for i, task := range *m.tasks {

        taskStr := task.Name
        if m.index == i {
            taskStr = m.styles.HighLightedTask.Render(taskStr)
        }
        taskStr += "\n"
        taskList += taskStr
    }

    if len(*m.tasks) > 0 {
        selectedTask :=  (*m.tasks)[m.index]
        infoBox = fmt.Sprintf("[*] %s\n[@] %s\n\n%s", selectedTask.Context, selectedTask.POI, selectedTask.Description)
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