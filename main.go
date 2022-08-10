package main

import (
	"net/http"

	"github.com/primaaku/project-go-prima/controllers/taskcontroller"
)

func main() {

	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task/get_form", taskcontroller.GetForm)
	http.HandleFunc("/task/store", taskcontroller.Store)
	http.HandleFunc("/task/delete", taskcontroller.Delete)
	http.HandleFunc("/task/complete", taskcontroller.Complete)

	http.ListenAndServe("", nil)

}
