package main

import (
	"net/http"

	"github.com/cehuda1/go-crud1/controllers/usercontrollers"
)

func main() {

	http.HandleFunc("/", usercontrollers.Index)
	http.HandleFunc("/user", usercontrollers.Index)
	http.HandleFunc("/user/index", usercontrollers.Index)
	http.HandleFunc("/user/add", usercontrollers.Add)
	http.HandleFunc("/user/Edit", usercontrollers.Edit)
	http.HandleFunc("/user/delete", usercontrollers.Delete)

	http.ListenAndServe(":4111", nil)
}