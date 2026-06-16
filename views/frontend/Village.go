package frontend

import (
	"fmt"
	"net/http"
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
