package taskcontroller

import (
	"bytes"
	"encoding/json"
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

func GetForm(w http.ResponseWriter, r *http.Request) {

	temp, _ := template.ParseFiles("views/task/form.html")
	temp.Execute(w, nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()
		var task entites.Task

		task.Assignee = r.Form.Get("assignee")
		task.Deadline = r.Form.Get("deadline")

		// id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
		err := taskModel.Create(&task)

		if err != nil {
			// insert data

			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		data := map[string]interface{}{
			"message": "Data berhasil disimpan",
			"data":    template.HTML(GetData()),
		}

		ResponseJson(w, http.StatusOK, data)
	}

}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
