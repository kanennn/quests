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

type Task struct {
	Name string `yaml:"name"`
	POI string `yaml:"poi"` //can be its own type later like "person" or "assignee"
	Context string `yaml:"context"` //can also be its own type later
	Description []byte
	// Precedence int `yaml:"precedence"`
	// Priority priority `yaml:"priority"`
}