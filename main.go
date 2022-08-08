package main

import (
	"net/http"

	"github.com/primaaku/project-go-prima/controllers/taskcontroller"
)

func main() {

	http.HandleFunc("/", taskcontroller.Index)

	http.ListenAndServe(":8000", nil)

}
