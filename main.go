package main

import (
	"net/http"

	"github.com/primaaku/go-project-prima/controllers/taskcontroller"
)

func main() {

	http.HandleFunc("/", taskcontroller.Index)

}
