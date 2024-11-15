package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type wait_model struct{}

func (m wait_model) Init() tea.Cmd                           { return nil }
func (m wait_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }
func (m wait_model) View() string                            { return "loading" }

type main_model struct {
	active_model tea.Model
	active_quest *quest
	styles       styles
	models       models
	width        int
	height       int
}

type models struct {
	lore_model     lore_model
	legend_model   legend_model
	children_model children_model
}

const (
	lore_id = iota
	legend_id
	children_id
)

type model_name struct {
	name string
	id   int
}

//todo how do we like, have pointers to sub and super quests without creating a recursive nightmare but that sort of preloads them
//todo mayhaps active quests load name, desc, files, logs, info, and subquests/superquests
//todo and then nonactive quests only load name and desc
//todo also what if quests just became like a wrapper for other softwares
//todo or like a bundle, like how git calls vi and vim calls shell actions
//* so maybe it's just like a viewer, and you call other things to change into other things lol
//* idk rly

func init_models(q *quest, s styles) models {
	lgm := legend_model{}
	lgm.field = textinput.New()
	lgm.quest = q

	lm := lore_model{}
	lm.field = textarea.New()
	lm.field.CharLimit = 0
	lm.quest = q

	cm := children_model{}
	cm.quest = q
	cm.s = s

	return models{
		lore_model:     lm,
		legend_model:   lgm,
		children_model: cm,
	}
}

func (m main_model) post_init() main_model {
	m.styles = default_styles()
	m.models = init_models(m.active_quest, m.styles)
	m.active_model = m.models.lore_model
	return m
}

func (m main_model) refill_models() main_model {
	switch am := m.active_model.(type) {
	case lore_model:
		m.models.lore_model = am
	case legend_model:
		m.models.legend_model = am
	case children_model:
		m.models.children_model = am
	}
	return m
}

func (m main_model) Init() tea.Cmd {
	return tea.Sequence(
		func() tea.Msg {
			q := new(quest)
			ex, err := os.Executable()
			Check(err)
			dir := filepath.Dir(ex)
			dir, _ = os.Getwd()
			q.peek(dir)
			q.open()
			return q
		},
		func() tea.Msg {
			return default_styles()
		},
	)
}

func (m main_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case (*quest):
		// should only be used once, before other models load
		// initial load
		m.active_quest = msg
		m = m.post_init()
	case quest:
		*m.active_quest = msg
		m.models = init_models(m.active_quest, m.styles)
		m.active_model = m.models.lore_model
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m = m.refill_models()
			m.active_model = m.models.lore_model
		case "2":
			m = m.refill_models()
			m.active_model = m.models.legend_model
		case "3":
			m = m.refill_models()
			m.active_model = m.models.children_model
		case "esc":
			if m.active_quest.parent != nil {
				return m, func() tea.Msg {
					c := m.active_quest.parent
					c.open() // i don't think this matters when the parent has already been opened. maybe we check for that // actually it helps to refresh because it rereads
					return *c
				}
			}
		case "ctrl+n":
			m.active_model = new_entry_model(m.active_quest)
		default:
			m.active_model, cmd = m.active_model.Update(msg)
			// case "enter":
			// 	i, ok := m.list.SelectedItem().(item)
			// 	if ok {
			// 		m.choice = string(i)
			// 	}
			// 	return m, te 3a.Quit
		}
	default:
		// fmt.Printf("unhandled message: %T", msg)
	}
	// m.active_model.Update(msg)
	return m, cmd
}

func (m main_model) View() string {
	var body string
	switch m.active_model.(type) {
	case *wait_model:
		body = m.active_model.View()
	default:
		body = header_view(m.active_quest, m.styles, m.active_model.View(), m.active_model)
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		m.styles.body.Render(body),
	)
}

func header_view(q *quest, s styles, content string, model tea.Model) string {
	model_names := []model_name{
		{name: "lore", id: lore_id},
		{name: "legend", id: legend_id},
		{name: "children", id: children_id},
	}
	var selected_id int
	switch model.(type) {
	case lore_model:
		selected_id = lore_id
	case legend_model:
		selected_id = legend_id
	case children_model:
		selected_id = children_id
	}
	var listings []string
	for _, n := range model_names {
		var model_listing string
		if n.id == selected_id {
			model_listing = s.menu_selected.Render(n.name)
		} else {
			model_listing = s.menu_unselected.Render(n.name)
		}
		listings = append(listings, model_listing)
	}

	header := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.PlaceHorizontal(
			s.inside_width,
			lipgloss.Center,
			s.quest_title.Render(q.Title),
			s.quest_title_whitespace...),
		lipgloss.PlaceHorizontal(
			s.inside_width,
			lipgloss.Center,
			s.quest_subtitle.Render(q.Subtitle),
		),
		s.menu_box.Render(lipgloss.JoinHorizontal(lipgloss.Top, listings...)),
	)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		content,
		s.footer.Render(q.dir),
	)
}

func tui() {
	m := new(main_model)
	m.active_model = new(wait_model)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there has been an error: %v", err)
		os.Exit(1)
	}
}
