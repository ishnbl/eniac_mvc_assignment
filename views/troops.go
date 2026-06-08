package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

type CreateTroopsReq struct {
	Type     string
	Quantity int
}

func CreateTroopsHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateTroopsReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(w, "Invalid Payload")
		return
	}
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}

	ret := models.CreateTroop(username, payload.Type, payload.Quantity)
	if ret != false {
		fmt.Fprintf(w, "Erro while creating troop")
		return
	}
	fmt.Fprintf(w, "Troop created successfully")
}

func MyTroopsHandler(w http.ResponseWriter, r *http.Request) {
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}

	myTroops := models.GetTroops(username)

	jsonResp, err := json.Marshal(myTroops)
	if err != nil {
		fmt.Fprintf(w, "Error creating JSON response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
