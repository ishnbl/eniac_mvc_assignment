package frontend

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

type IndBuildRes struct {
	Type    string
	LevSpec []models.LevelSpecific
	Yield   int
}

type IndBuildSto struct {
	Type    string
	LevSpec []models.LevelSpecific
	Storage int
}
type ReturnShop struct {
	Troops    []models.Troops
	Defenses  []models.Defenses
	ResBuilds []models.RetShopBuilding
}

func ShopHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		fmt.Fprintf(w, "you need to login first")
		return
	}

	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/web/login", http.StatusSeeOther)
		return
	}
	var retSho ReturnShop

	troops := models.GetShopTroops(username)
	defenses := models.GetShopDefenses(username)
	buildings := models.GetShopBuildings(username)

	retSho = ReturnShop{
		Troops:    troops,
		Defenses:  defenses,
		ResBuilds: buildings,
	}
	tmpl := template.Must(template.ParseGlob("views/templates/*.html"))
	err = tmpl.ExecuteTemplate(w, "shop.html", retSho)
	if err != nil {
		fmt.Fprintf(w, "could not render the shop page, something went wrong")
	}

}

func BuyTroopHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "couldnt read the form data")
		return
	}

	troopType := r.FormValue("troop_type")
	level, err := strconv.Atoi(r.FormValue("troop_level"))
	quantity, e := strconv.Atoi(r.FormValue("quantity"))
	if troopType == "" || err != nil || e != nil || quantity < 1 {
		fmt.Fprintf(w, "invalid troop selection, check the type level and quantity")
		return
	}

	ok, reason := models.CreateTroop(username, troopType, quantity, level)
	if !ok {
		fmt.Fprintf(w, "couldnt buy troops: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/shop")
	w.WriteHeader(http.StatusOK)
}

func BuyDefenseHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "couldnt read the form data")
		return
	}

	defenseType := r.FormValue("defense_type")
	am, err := strconv.Atoi(r.FormValue("amount"))
	if defenseType == "" || err != nil || am < 1 {
		fmt.Fprintf(w, "invalid defense selection, check the type and amount")
		return
	}

	ok, reason := models.CreateDefense(username, defenseType, am)
	if !ok {
		fmt.Fprintf(w, "couldnt buy defense: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/shop")
	w.WriteHeader(http.StatusOK)
}

func BuyBuildingHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "couldnt read the form data")
		return
	}

	buildingType := r.FormValue("building_type")
	x, err := strconv.Atoi(r.FormValue("x"))
	y, e := strconv.Atoi(r.FormValue("y"))
	if buildingType == "" || err != nil || e != nil {
		fmt.Fprintf(w, "invalid building selection, need a type and valid x y coords")
		return
	}

	ok, reason := models.CreateBuilding(username, buildingType, x, y)
	if !ok {
		fmt.Fprintf(w, "couldnt place building: %s", reason)
		return
	}

	w.Header().Set("HX-Redirect", "/web/app/shop")
	w.WriteHeader(http.StatusOK)
}
