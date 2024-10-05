package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type wait_model struct{}

func (m wait_model) Init() tea.Cmd                           { return nil }
func (m wait_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }
func (m wait_model) View() string                            { return "loading" }

type main_model struct {
	active_model tea.Model
	active_quest *quest
	models       struct {
		info_model     *info_model
		legend_model   *legend_model
		children_model *children_model
	}
}

type model_load struct{ m tea.Model }

//todo how do we like, have pointers to sub and super quests without creating a recursive nightmare but that sort of preloads them
//todo mayhaps active quests load name, desc, files, logs, info, and subquests/superquests
//todo and then nonactive quests only load name and desc
//todo also what if quests just became like a wrapper for other softwares
//todo or like a bundle, like how git calls vi and vim calls shell actions
//* so maybe it's just like a viewer, and you call other things to change into other things lol
//* idk rly

func (m main_model) Init() tea.Cmd {
	return tea.Sequence(func() tea.Msg {
		q := new(quest)
		q.open("")
		return q
	}, func() tea.Msg {
		return model_load{m: new(info_model)}
	},
		func() tea.Msg {
			return model_load{m: new(legend_model)}
		}, func() tea.Msg {
			return model_load{m: new(children_model)}
		},
	)
}

func (m main_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case (*quest):
		m.active_quest = msg
	case model_load:
		switch msg.m.(type) {
		case *info_model:
			m.active_model = msg.m
			m.models.info_model = msg.m.(*info_model)
			(m.active_model.(*info_model)).quest = m.active_quest
		case *legend_model:
			(msg.m.(*legend_model)).quest = m.active_quest
			m.models.legend_model = msg.m.(*legend_model)
		case *children_model:
			(msg.m.(*children_model)).quest = m.active_quest
			m.models.children_model = msg.m.(*children_model)
		}
	case tea.Model:
		m.active_model = msg
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			cmd = func() tea.Msg { return m.models.info_model } //this will break if the model is not loaded yet
		case "2":
			cmd = func() tea.Msg { return m.models.legend_model } //this will break if the model is not loaded yet
		case "3":
			cmd = func() tea.Msg { return m.models.children_model } //this will break if the model is not loaded yet
			// case "enter":
			// 	i, ok := m.list.SelectedItem().(item)
			// 	if ok {
			// 		m.choice = string(i)
			// 	}
			// 	return m, tea.Quit
		}
	default:
		//fmt.Printf("unhandled message: %T", msg)
	}
	m.active_model.Update(msg)
	return m, cmd
}

func (m main_model) View() string {
	if m.active_quest != nil {

		head := m.active_quest.Name
		active_view := m.active_model.View()
		return head + "\n" + active_view
	} else {
		return "aaAA"
	}
}

func tui() {
	m := new(main_model)
	m.active_model = new(wait_model)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there has been an error: %v", err)
		os.Exit(1)
	}
}