package frontend

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

type VillData struct {
	VillageLevel int
	Gold         int
	Oil          int
	FarmLand     int
	Money        int
	UpgradeCost  int
	Buildings    []models.RetBuilding
	Troops       []models.RetTroops
	Defenses     []models.RetDefenses
}

func Village(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Redirect(w, r, "/web/login", http.StatusSeeOther)
		return
	}

	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/web/login", http.StatusSeeOther)
		return
	}

	villData := models.GetVillage(username)

	troops := models.GetTroops(username)
	defenses := models.GetDefenses(username)
	buildings := models.GetBuildings(username)

	viewData := VillData{
		VillageLevel: villData.VillageLevel,
		Gold:         villData.Gold,
		Oil:          villData.Oil,
		FarmLand:     villData.FarmLand,
		Money:        villData.Money,
		UpgradeCost:  villData.VillageLevel * 1000,
		Buildings:    buildings,
		Troops:       troops,
		Defenses:     defenses,
	}

	tmpl := template.Must(template.ParseGlob("views/templates/*.html"))
	err = tmpl.ExecuteTemplate(w, "village.html", viewData)
	if err != nil {
		fmt.Fprintf(w, "could not render template")
	}
}

func MoveBuildingPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Redirect(w, r, "/web/login", http.StatusSeeOther)
		return
	}
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/web/login", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("buildingID"))
	if err != nil {
		fmt.Fprintf(w, "invalid building id")
		return
	}

	buildings := models.GetBuildings(username)
	var found models.RetBuilding
	found_hit := false
	for i := range buildings {
		if buildings[i].ID == uint(id) {
			found = buildings[i]
			found_hit = true
			break
		}
	}
	if !found_hit {
		fmt.Fprintf(w, "building not found")
		return
	}

	tmpl := template.Must(template.ParseGlob("views/templates/*.html"))
	err = tmpl.ExecuteTemplate(w, "move_building.html", found)
	if err != nil {
		fmt.Fprintf(w, "could not render template")
	}
}

func MoveBuildingHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "couldnt read the form data")
		return
	}

	id, err := strconv.Atoi(r.FormValue("buildingID"))
	x, e1 := strconv.Atoi(r.FormValue("x"))
	y, e2 := strconv.Atoi(r.FormValue("y"))
	if err != nil || e1 != nil || e2 != nil {
		fmt.Fprintf(w, "invalid input")
		return
	}

	ok, reason := models.MoveBuilding(username, uint(id), x, y)
	if !ok {
		fmt.Fprintf(w, "couldnt move building: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/village")
	w.WriteHeader(http.StatusOK)
}
