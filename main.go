package main

import (
	"net/http"
	"os"

	"github.com/primaaku/project-go-prima/controllers/taskcontroller"
)

func main() {

	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task/get_form", taskcontroller.GetForm)
	http.HandleFunc("/task/store", taskcontroller.Store)
	http.HandleFunc("/task/delete", taskcontroller.Delete)
	http.HandleFunc("/task/complete", taskcontroller.Complete)

	// http.ListenAndServe(":8000", nil)

	// appPort := ":" + os.Getenv("PORT")
	// fmt.Println(appPort) // This prints ":8000"
	// s := &http.Server{
	// 	Addr: ":8000",
	// }
	http.ListenAndServe(os.Getenv("PORT"), nil)

}
