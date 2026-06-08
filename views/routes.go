package views

import (
	"github.com/gorilla/mux"
	"github.com/ishnbl/eniac_mvc_assignment/controller"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(controller.AuthMiddleware)
	protected.HandleFunc("/village", VillageHandler).Methods("GET")
	protected.HandleFunc("/buildings", CreateBuildingsHandler).Methods("POST")
	protected.HandleFunc("/defenses", CreateDefensesHandler).Methods("POST")
	protected.HandleFunc("/troops", CreateTroopsHandler).Methods("POST")
	protected.HandleFunc("/mytroops", MyTroopsHandler).Methods("GET")
	protected.HandleFunc("/mydefenses", MyDefensesHandler).Methods("GET")
	protected.HandleFunc("/shop/troops", ShopTroopsHandler).Methods("GET")
	protected.HandleFunc("/shop/defenses", ShopDefensesHandler).Methods("GET")
}
