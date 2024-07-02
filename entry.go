package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type entry struct {
    task Task
    inputField textinput.Model
    styles *Styles
    fields []field
    index int
}

type field struct {
    question string
    fill func (t *Task, s string)
}

func newEntry() *entry {
    entry := new(entry)

    entry.inputField = textinput.New()
    entry.inputField.Focus()

    entry.fields = []field{
        {question: "Name", fill: (*Task).fillName},
        {question: "POI", fill: (*Task).fillPOI},
        {question: "Context", fill: (*Task).fillContext},
        {question: "Description", fill: (*Task).fillDescription},
}

    return entry
}

func (m entry) Init() tea.Cmd {
    return nil
}

func (m entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    
    switch msg := msg.(type) {

        case tea.KeyMsg:    
            switch msg.String() {

            case "enter":
                // TODO we will need an input handler for non-string task fields later

                field := m.fields[m.index]
                field.fill(&m.task, m.inputField.Value())
                
                if m.index < len(m.fields)-1 {
                    m.index++
                    m.inputField.Reset()
                } else {
                    return m, func () tea.Msg {return entryReturn{t: m.task}}
                }
            }
    }
    m.inputField, cmd = m.inputField.Update(msg)
    return m, cmd
}

func (m entry) View() string {
    m.inputField.Placeholder = m.fields[m.index].question
    return m.styles.InputField.Render(m.inputField.View())
}