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
		fmt.Fprintf(w, "bad req")
		return
	}

	create := models.UpgradeVillage(username)
	if create == false {
		fmt.Fprintf(w, "could not create trop")
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
		fmt.Fprintf(w, "bad req")
		return
	}

	id, err := strconv.Atoi(r.FormValue("buildingID"))
	if err != nil {
		fmt.Fprintf(w, "invalid building")
		return
	}

	create := models.UpgradeBuilding(username, uint(id))
	if create == false {
		fmt.Fprintf(w, "could not create trop")
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
		fmt.Fprintf(w, "bad req")
		return
	}

	id, err := strconv.Atoi(r.FormValue("buildingID"))
	if err != nil {
		fmt.Fprintf(w, "invalid building")
		return
	}

	create := models.ExtractResources(username, uint(id))
	if create == false {
		fmt.Fprintf(w, "could not create trop")
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/village")
	w.WriteHeader(http.StatusOK)
}
