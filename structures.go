package main

import "fmt"

type Task struct { //TODO this could be similar to a python dictionary if it would make things simpler in the future
	Name string `yaml:"name"`
	POI string `yaml:"poi"` //can be its own type later like "person" or "assignee"
	Context string `yaml:"context"` //can also be its own type later
	Description string // TODO we will need a input handler in the tui for not string inputs later
	Completed bool 
	Hidden bool //this is used for soft deleting, basically hiding tasks //needs to be "deactive"
	Class class
	// Precedence int `yaml:"precedence"`
	// Priority priority `yaml:"priority"`
}

type class int

const ( //this could use better names, like journey, adventure, hike, joust, hunt, etc
	main_quest = iota
	side_quest
	mini_quest
	sleeping_quest
)

const classLen = 4


// These are import functions: they take in a string, and map that into the correct type for the struct field
// TODO We may eventually need export functions for mapping struct fields into strings
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


func (t *Task) fillCompleted(s string) {
	if s == "true" {
		t.Completed = true
	} else if s == "false" {
		t.Completed = false
	} else {
		panic(fmt.Sprintf(`Invalid completion string "%s"`, s))
	}
}
func (t *Task) fillHidden(s string) {
	if s == "true" {
		t.Hidden = true
	} else if s == "false" {
		t.Hidden = false
	} else {
		panic(fmt.Sprintf(`Invalid hidden state string "%s"`, s))
	}
}

func (t *Task) fillClass(s string) {
	switch s {
	case "main_quest": t.Class = main_quest
	case "side_quest": t.Class = side_quest
	case "mini_quest": t.Class = mini_quest
	case "sleeping_quest": t.Class = sleeping_quest
	default: panic(fmt.Sprintf(`Invalid class type "%s"`, s))
	}
}

// export functions don't neccesarily need to be methods of pointers (*task)
func (t *Task) exportName() string {return t.Name}
func (t *Task) exportPOI() string {return t.POI}
func (t *Task) exportContext() string {return t.Context}
func (t *Task) exportDescription() string {return t.Description}

func (t *Task) exportCompleted() string {
	if t.Completed {
		return "true"
	} else {
		return "false"
	}
}
func (t *Task) exportHidden() string {
	if t.Hidden {
		return "true"
	} else {
		return "false"
	}
}

func (t *Task) exportClass() string {
	switch t.Class {
		case main_quest: return "main_quest"
		case side_quest: return "side_quest"
		case mini_quest: return "mini_quest"
		case sleeping_quest: return "sleeping_quest"
	default: panic("What the heck is class type" + string(t.Class))
	}
}