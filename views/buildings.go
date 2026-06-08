package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

type CreateBuildingReq struct {
	X    int
	Y    int
	Type string
}

func CreateBuildingsHandler(w http.ResponseWriter, r *http.Request) {
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}
	var payload CreateBuildingReq
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(w, "Invalid Payload")
		return
	}

	ret := models.CreateBuilding(username, payload.Type, payload.X, payload.Y)
	if ret != false {
		fmt.Fprintf(w, "Error while creating building")
		return
	}
	fmt.Fprintf(w, "Building created successfully")
}
