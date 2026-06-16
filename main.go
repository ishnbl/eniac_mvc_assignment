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
	db := models.DB
	db.AutoMigrate(&models.User{}, &models.Village{}, &models.Defenses{}, &models.VillageDefMapping{}, &models.Troops{}, &models.TroopVillageMapping{}, &models.Building{}, &models.LevelSpecific{}, &models.ResourceColl{}, &models.Storage{}, &models.BattleReplay{}, &models.ReplayDefenses{}, &models.ReplayTroops{})
	server.ListenAndServe()
}
