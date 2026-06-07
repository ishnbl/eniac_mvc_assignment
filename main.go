package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"github.com/ishnbl/eniac_mvc_assignment/views"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", views.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", views.LoginHandler).Methods("POST")
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(controller.AuthMiddleware)
	protected.HandleFunc("/village", views.VillageHandler).Methods("GET")
	protected.HandleFunc("/buildings", views.CreateBuildingsHandler).Methods("POST")
	protected.HandleFunc("/defenses", views.CreateDefensesHandler).Methods("POST")
	protected.HandleFunc("/troops", views.CreateTroopsHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	fmt.Println("Server is running on port 8000")
	fmt.Println("Hello, MVC!")
	models.Init_DB()
	db := models.DB
	db.AutoMigrate(&models.User{}, &models.Village{}, &models.Defenses{}, &models.VillageDefMapping{}, &models.Troops{}, &models.TroopVillageMapping{}, &models.Buildings{}, &models.BuildingsVillageMapping{}, &models.BattleReplay{}, &models.ReplayDefenses{}, &models.ReplayTroops{})
	server.ListenAndServe()
}
