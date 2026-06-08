package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

type CreateDefensesReq struct {
	Type   string
	Amount int
}

func CreateDefensesHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateDefensesReq
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

	ret := models.CreateDefense(username, payload.Type, payload.Amount)
	if ret != false {
		fmt.Fprintf(w, "Error while creating")
		return
	}
	fmt.Fprintf(w, "Defense created successfully")
}

func MyDefensesHandler(w http.ResponseWriter, r *http.Request) {
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}

	myDefenses := models.GetDefenses(username)

	jsonResp, err := json.Marshal(myDefenses)
	if err != nil {
		fmt.Fprintf(w, "Error creating JSON response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
