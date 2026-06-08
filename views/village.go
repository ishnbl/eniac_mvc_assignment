package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

func VillageHandler(w http.ResponseWriter, r *http.Request) {
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}

	VillageRet := models.GetVillageBuildings(username)
	jsonResp, err := json.Marshal(VillageRet)
	if err != nil {
		fmt.Fprintf(w, "Error creating JSON response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
