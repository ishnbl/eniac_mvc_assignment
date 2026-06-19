package views

import (
	"github.com/gorilla/mux"
	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/views/frontend"
)

func SetupRoutes(r *mux.Router) {
	//NEW ENDPOINTS WORK WITH HTMX
	front := r.PathPrefix("/web/app").Subrouter()
	front.Use(controller.AuthMiddlewareHtmx)
	r.HandleFunc("/web/login", frontend.LoginView).Methods("GET")
	r.HandleFunc("/web/login", frontend.LoginViewHtmx).Methods("POST")
	r.HandleFunc("/web/register", frontend.RegisterHandlerHtmx).Methods("POST")
	front.HandleFunc("/shop", frontend.ShopHandlerHtmx).Methods("GET")
	front.HandleFunc("/village", frontend.Village).Methods("GET")
	front.HandleFunc("/battles", frontend.BattlesHandlerHtmx).Methods("GET")
	front.HandleFunc("/buy/troops", frontend.BuyTroopHandlerHtmx).Methods("POST")
	front.HandleFunc("/buy/defenses", frontend.BuyDefenseHandlerHtmx).Methods("POST")
	front.HandleFunc("/buy/buildings", frontend.BuyBuildingHandlerHtmx).Methods("POST")
	front.HandleFunc("/upgrade/village", frontend.UpgradeVillageHandlerHtmx).Methods("POST")
	front.HandleFunc("/upgrade/building", frontend.UpgradeBuildingHandlerHtmx).Methods("POST")
	front.HandleFunc("/collect", frontend.ExtractHandlerHtmx).Methods("POST")
	front.HandleFunc("/battle/{villageID}", frontend.DoBattleHandlerHtmx).Methods("GET")
	front.HandleFunc("/fight", frontend.FightHandlerHtmx).Methods("POST")
	front.HandleFunc("/move/building", frontend.MoveBuildingPageHandler).Methods("GET")
	front.HandleFunc("/move/building", frontend.MoveBuildingHandler).Methods("POST")

}
