package main

// hi this is my io package named leo

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ReadFiles() [][]byte {
	dirs, err := os.ReadDir(filepath.Join("quests"))
	Check(err)
	var files [][]byte 
	for _, dir := range dirs {
		if !dir.IsDir() {
			file, err := os.ReadFile(filepath.Join("quests", dir.Name()))
			Check(err)
			files = append(files, file)
		}
	}

	return files
}

// includes description
type importMap struct {
	regex string
	fill func (t *Task, s string)
}

// does not include description
type exportMap struct {
	yamlName string
	export func (t *Task) (s string)
}



func ParseFiles(files [][]byte) (tasks []Task) {
	//I need to just make a whole fucnction that splits the files into frontmatter and content, and not pass the whole contents around to each step in parsing

	yamlParser := `---\n[\S\s]*%s:\s*([\S\s]+?)\n[\S\s]*---` //if i want only words and no symbols, i need to change this and change what items are allowed on entry
	descParser := `---[\S\s]+?---\s*([\S\s]*)`

	importer := []importMap{
		{regex: fmt.Sprintf(yamlParser, "name"), fill: (*Task).fillName},
		{regex: fmt.Sprintf(yamlParser, "poi"), fill: (*Task).fillPOI},
		{regex: fmt.Sprintf(yamlParser, "context"), fill: (*Task).fillContext},
		{regex: fmt.Sprintf(yamlParser, "completed"), fill: (*Task).fillCompleted},
		{regex: fmt.Sprintf(yamlParser, "hidden"), fill: (*Task).fillHidden},
		{regex: descParser, fill: (*Task).fillDescription},
	}

	for _, file := range files {

		

		task := new(Task)

		for _, v := range importer {
			re, err := regexp.Compile(v.regex)
			Check(err)
			v.fill(task, string(re.FindSubmatch(file)[1])) //TODO maybe throw []byte to imports? just depends on where i want to convert []byte to string and vice versa
		}

		tasks = append(tasks, *task)	
	}
	return tasks
}

func CreateTasks(tasks []Task) {

	exporter := []exportMap{
		{yamlName: "name", export: (*Task).exportName},
		{yamlName: "poi", export: (*Task).exportPOI},
		{yamlName: "context", export: (*Task).exportContext},
		{yamlName: "completed", export: (*Task).exportCompleted},
		{yamlName: "hidden", export: (*Task).exportHidden},
	}
	
	for _, task := range tasks {
		//yaml := MakeYamsFromTask(task)
		var yaml string
		for _, v := range exporter {
			yaml += fmt.Sprintf("%s: %s\n", v.yamlName, v.export(&task)) // this may not need to be a pointer
		}

		data := fmt.Sprintf("---\n%s---\n\n%s", yaml, task.exportDescription())
		err := os.WriteFile(filepath.Join("quests", task.Name + ".md"), []byte(data), 0600 )
		Check(err)

	}
	
}

func getTasks() []Task {
    files := ReadFiles()
	tasks := ParseFiles(files)
    return tasks
}