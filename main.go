package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"github.com/ishnbl/eniac_mvc_assignment/views"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", views.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", views.LoginHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	fmt.Println("Server is running on port 8000")
	fmt.Println("Hello, MVC!")
	models.Init_DB()
	db := models.DB
	db.AutoMigrate(&models.User{}, &models.Village{}, &models.Defenses{}, &models.BattleReplay{}, &models.Troops{}, &models.Buildings{}, &models.ReplayDefenses{}, &models.ReplayTroops{})
	server.ListenAndServe()
}
