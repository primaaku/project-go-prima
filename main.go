package main

import (
	"net/http"

	"github.com/primaaku/project-go-prima/controllers/taskcontroller"
)

func main() {

	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task/get_form", taskcontroller.GetForm)
	http.HandleFunc("/task/store", taskcontroller.Store)

	http.ListenAndServe(":8000", nil)

}
