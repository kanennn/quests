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

type parseMap struct {
	regex string
	fill func (t *Task, s string)
}

func ParseFiles(files [][]byte) (tasks []Task) {
	//I need to just make a whole fucnction that splits the files into frontmatter and content, and not pass the whole contents around to each step in parsing

	yamlParser := `---\n[\S\s]*%s:\s*([\S\s]+?)\n[\S\s]*---` //if i want only words and no symbols, i need to change this and change what items are allowed on entry
	descParser := `---[\S\s]+?---\s*([\S\s]*)`

	parser := []parseMap{
		{regex: fmt.Sprintf(yamlParser, "name"), fill: (*Task).fillName},
		{regex: fmt.Sprintf(yamlParser, "poi"), fill: (*Task).fillPOI},
		{regex: fmt.Sprintf(yamlParser, "context"), fill: (*Task).fillContext},
		{regex: descParser, fill: (*Task).fillDescription},
	}

	for _, file := range files {

		

		task := new(Task)

		for _, v := range parser {
			re, err := regexp.Compile(v.regex)
			Check(err)

			fmt.Println("Regex:")
			fmt.Println(v.regex)
			fmt.Print("submatch:")
			fmt.Println(string(re.FindSubmatch(file)[1]))
			v.fill(task, string(re.FindSubmatch(file)[1])) //TODO maybe throw []byte to imports? just depends on where i want to convert []byte to string and vice versa
		}

		tasks = append(tasks, *task)	
	}
	return tasks
}

func CreateTasks(tasks []Task) {
	
	for _, task := range tasks {
		yaml := MakeYamsFromTask(task)
		data := fmt.Sprintf("---\n%s\n---\n\n%s", yaml, task.Description)
		err := os.WriteFile(filepath.Join("quests", task.Name + ".md"), []byte(data), 0600 )
		Check(err)

	}
	
}

func getTasks() []Task {
    files := ReadFiles()
	tasks := ParseFiles(files)
    return tasks
}