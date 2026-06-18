package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"github.com/ishnbl/eniac_mvc_assignment/views"
)

func main() {

	r := mux.NewRouter()
	views.SetupRoutes(r)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	fmt.Println("Server is running on port 8000")
	fmt.Println("Hello, MVC!")
	models.Init_DB()
	server.ListenAndServe()
}
