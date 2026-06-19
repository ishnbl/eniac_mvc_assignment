package frontend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

func UpgradeVillageHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "couldnt parse the form, something is wrong with the request")
		return
	}

	ok, reason := models.UpgradeVillage(username)
	if !ok {
		fmt.Fprintf(w, "upgrade failed: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/village")
	w.WriteHeader(http.StatusOK)
}

func UpgradeBuildingHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "couldnt parse the form, something is wrong with the request")
		return
	}

	id, err := strconv.Atoi(r.FormValue("buildingID"))
	if err != nil {
		fmt.Fprintf(w, "invalid building id, make sure its a number")
		return
	}

	ok, reason := models.UpgradeBuilding(username, uint(id))
	if !ok {
		fmt.Fprintf(w, "upgrade failed: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/village")
	w.WriteHeader(http.StatusOK)
}

func ExtractHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "couldnt parse the form, something is wrong with the request")
		return
	}

	id, err := strconv.Atoi(r.FormValue("buildingID"))
	if err != nil {
		fmt.Fprintf(w, "invalid building id, make sure its a number")
		return
	}

	ok, reason := models.ExtractResources(username, uint(id))
	if !ok {
		fmt.Fprintf(w, "couldnt collect resources: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/village")
	w.WriteHeader(http.StatusOK)
}
