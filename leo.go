package main

// hi this is my io package named leo

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ReadFiles() [][]byte {
	dirs, err := os.ReadDir(filepath.Join("playground", "input"))
	var files [][]byte 
	Check(err)
	for _, dir := range dirs {
		if !dir.IsDir() {
			file, err := os.ReadFile(filepath.Join("playground", "input", dir.Name()))
			Check(err)
			files = append(files, file)
		}
	}

	return files
}

func ParseFiles(files [][]byte) (tasks []Task) {
	//I need to just make a whole fucnction that splits the files into frontmatter and content, and not pass the whole contents around to each step in parsing
	for _, file := range files {
		task := ReadYamsFromFrontatter(file)
		re, err := regexp.Compile(`---[\S\s]+?---\s*([\S\s]*)`)
		Check(err)
		desc := re.FindSubmatch(file)[1]
		task.Description = string(desc) // TODO may be needed to switch back to byte later, when the task struct is not just strings
		tasks = append(tasks, task)
	}
	return tasks
}

func CreateTasks(tasks []Task) {
	
	for _, task := range tasks {
		yaml := MakeYamsFromTask(task)
		data := fmt.Sprintf("---\n%s\n---\n\n%s", yaml, task.Description)
		err := os.WriteFile(filepath.Join("playground", "output", task.Name + ".md"), []byte(data), 0600 )
		Check(err)

	}
	
}

func getTasks() []Task {
    files := ReadFiles()
	tasks := ParseFiles(files)
    return tasks
}