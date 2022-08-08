package taskcontroller

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/primaaku/project-go-prima/entites"
	"github.com/primaaku/project-go-prima/models/taskmodel"
)

var taskModel = taskmodel.New()

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Berjalan")
	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/task/index.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}

	// func registered untuk view
	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/task/data.html")

	var task []entites.Task
	err := taskModel.FindAll(&task)
	if err != nil {
		panic(err)
	}

	fmt.Println(task)
	// func mengirim data task
	data := map[string]interface{}{
		"task": task,
	}

	// ambil template
	temp.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()
}
