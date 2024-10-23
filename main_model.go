package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

type wait_model struct{}

func (m wait_model) Init() tea.Cmd                           { return nil }
func (m wait_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }
func (m wait_model) View() string                            { return "loading" }

type main_model struct {
	active_model tea.Model
	active_quest *quest
	models       *models
}

type models struct {
	info_model     info_model
	legend_model   legend_model
	children_model children_model
	lore_model     lore_model
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
	return tea.Sequence(
		func() tea.Msg {
			return new(models)
		},
		func() tea.Msg {
			q := new(quest)
			ex, err := os.Executable()
			Check(err)
			dir := filepath.Dir(ex)
			dir, _ = os.Getwd()
			q.peek(dir)
			q.open()
			return q
		}, func() tea.Msg {
			return model_load{m: new(info_model)}
		},
		func() tea.Msg {
			return model_load{m: new(legend_model)}
		}, func() tea.Msg {
			return model_load{m: new(children_model)}
		},
		func() tea.Msg {
			return model_load{m: new(lore_model)}
		},
	)
}

func (m main_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case *models:
		m.models = msg
	case (*quest):
		// should only be used once, before other models load
		// initial load
		m.active_quest = msg

		if msg.parent != nil && (msg.Name == msg.parent.Name || msg.Description == msg.parent.Description) {
			panic(msg.Name + " " + msg.parent.Name)
		}
	case quest:
		*m.active_quest = msg
		m.active_model = &m.models.info_model
	case model_load:
		switch msg.m.(type) {
		case *info_model:
			(msg.m.(*info_model)).quest = m.active_quest
			m.models.info_model = *(msg.m.(*info_model))
			m.active_model = msg.m
		case *legend_model:
			(msg.m.(*legend_model)).quest = m.active_quest
			m.models.legend_model = *msg.m.(*legend_model)
		case *children_model:
			(msg.m.(*children_model)).quest = m.active_quest
			(msg.m.(*children_model)).models = m.models
			m.models.children_model = *msg.m.(*children_model)
		case *lore_model:
			(msg.m.(*lore_model)).quest = m.active_quest
			m.models.lore_model = *msg.m.(*lore_model)
		}
	case tea.Model:
		m.active_model = msg
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit
		case "1":
			cmd = func() tea.Msg { return &m.models.info_model } // this will break if the model is not loaded yet
		case "2":
			cmd = func() tea.Msg { return &m.models.legend_model } // this will break if the model is not loaded yet
		case "3":
			cmd = func() tea.Msg { return &m.models.children_model } // this will break if the model is not loaded yet
		case "4":
			cmd = func() tea.Msg { return &m.models.lore_model } // this will break if the model is not loaded yet
		case "esc":
			if m.active_quest.parent != nil {
				return m, func() tea.Msg {
					c := m.active_quest.parent
					c.open() // i don't think this matters when the parent has already been opened. maybe we check for that // actually it helps to refresh because it rereads
					return *c
				}
			}
		case "ctrl+n":
			return m, func() tea.Msg { return new_entry_model(m.active_quest) }
		default:
			m.active_model, cmd = m.active_model.Update(msg)

			// case "enter":
			// 	i, ok := m.list.SelectedItem().(item)
			// 	if ok {
			// 		m.choice = string(i)
			// 	}
			// 	return m, tea.Quit
		}
	default:
		// fmt.Printf("unhandled message: %T", msg)
	}
	// m.active_model.Update(msg)
	return m, cmd
}

func (m main_model) View() string {
	if m.active_quest != nil {

		var view string
		switch m.active_model.(type) {
		case *info_model:
			view = "info"
		case *legend_model:
			view = "legend_model"
		case *children_model:
			view = "children"
		case *lore_model:
			view = "lore"
		case *entry_model:
			view = "new"
		}

		head := m.active_quest.Name
		active_view := m.active_model.View()
		return view + "@" + head + "\n\n" + active_view
	} else {
		return "aaAA"
	}
}

func tui() {
	m := new(main_model)
	m.active_model = new(wait_model)
	p := tea.NewProgram(m) // tea.WithAltScreen()
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there has been an error: %v", err)
		os.Exit(1)
	}
}
