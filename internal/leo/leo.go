package leo

// hi this is my io package named leo

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kanennn/quests/internal/structures"
	"github.com/kanennn/quests/internal/yaml"
	"github.com/kanennn/quests/util"
)

func ReadFiles() [][]byte {
	dirs, err := os.ReadDir(filepath.Join("playground", "input"))
	var files [][]byte 
	util.Check(err)
	for _, dir := range dirs {
		if !dir.IsDir() {
			file, err := os.ReadFile(filepath.Join("playground", "input", dir.Name()))
			util.Check(err)
			files = append(files, file)
		}
	}

	return files
}

func ParseFiles(files [][]byte) (tasks []structures.Task) {
	//I need to just make a whole fucnction that splits the files into frontmatter and content, and not pass the whole contents around to each step in parsing
	for _, file := range files {
		task := yaml.ReadYamsFromFrontatter(file)
		re, err := regexp.Compile(`---[\S\s]+?---\s*([\S\s]*)`)
		util.Check(err)
		desc := re.FindSubmatch(file)[1]
		task.Description = desc
		tasks = append(tasks, task)
	}
	return tasks
}

func CreateTasks(tasks []structures.Task) {
	
	for _, task := range tasks {
		yaml := yaml.MakeYamsFromTask(task)
		data := fmt.Sprintf("---\n%s\n---\n\n%s", yaml, task.Description)
		err := os.WriteFile(filepath.Join("playground", "output", task.Name + ".md"), []byte(data), 0600 )
		util.Check(err)

	}
	
}