package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type value interface{}

type field interface {
	Update(tea.Msg) (field, tea.Cmd)
	View() string
	Value() value
	Focus()
	Reset()
}

type text_field struct {
	model textinput.Model
}

func new_text_field() text_field {
	m := textinput.New()
	m.Focus()
	return text_field{
		model: m,
	}
}

func (f text_field) Value() value {
	return f.model.Value()
}
func (f text_field) Reset() {
	f.model.Reset()
}
func (f text_field) Focus() {
	f.model.Focus()
}

func (f text_field) Update(msg tea.Msg) (field, tea.Cmd) {
	var cmd tea.Cmd
	f.model, cmd = f.model.Update(msg)
	return f, cmd
}

func (f text_field) View() string {
	return f.model.View()
}
