package taskcontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Berjalan")
	temp, _ := template.ParseFiles("views/task/index.html")
	temp.Execute(w, nil)
}
