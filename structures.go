package main

type priority struct {
	name string
	description string
}

var ( 
	High = priority{
		name: "High",
		description: "aaaa",
	}
	Medium = priority{
		name: "Medium",
		description: "aaaa",
	}
	Low = priority{
		name: "Medium",
		description: "aaaa",
	}
	)

type Task struct { //TODO this could be similar to a python dictionary if it would make things simpler in the future
	Name string `yaml:"name"`
	POI string `yaml:"poi"` //can be its own type later like "person" or "assignee"
	Context string `yaml:"context"` //can also be its own type later
	Description string // TODO we will need a input handler in the tui for not string inputs later
	// Precedence int `yaml:"precedence"`
	// Priority priority `yaml:"priority"`
} 

func (t *Task) fillName(s string) {
	t.Name = s
}
func (t *Task) fillPOI(s string) {
	t.POI = s
}
func (t *Task) fillContext(s string) {
	t.Context = s
}
func (t *Task) fillDescription(s string) {
	t.Description = s
}