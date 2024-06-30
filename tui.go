package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    tasks []Task                // task structs loaded from local files
    cursor   int                // which task item our cursor is pointing at
}

type Tasks []Task

func readTasks() tea.Msg {
	files := ReadFiles()
	tasks := ParseFiles(files)
	return Tasks(tasks)
}

func (m model) Init() (tea.Cmd) {
	return readTasks
}



func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:


        // Cool, what was the actual key pressed?
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.tasks)-1 {
                m.cursor++
            }
        }
	
	case Tasks:
		m.tasks = msg
	}

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {

    s := "Welcome to Quests"
    s += "\nPress q to quit.\n\n"

    // Iterate over our choices
    for i, task := range m.tasks {
        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, task.Name)
    }

    if len(m.tasks) > 0 {
        selectedTask := m.tasks[m.cursor]
        s += fmt.Sprintf("\n[*] %s\n[@] %s\n\n%s\n", selectedTask.Context, selectedTask.POI, selectedTask.Description)
    }

    

    // Send the UI for rendering
    return s
}

func tui() {
    p := tea.NewProgram(model{})
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}