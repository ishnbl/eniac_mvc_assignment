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
		fmt.Fprintf(w, "Login First")
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
		fmt.Fprintf(w, "could not render template")
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
		fmt.Fprintf(w, "bad req")
		return
	}

	troopType := r.FormValue("troop_type")
	level, err := strconv.Atoi(r.FormValue("troop_level"))
	quantity, e := strconv.Atoi(r.FormValue("quantity"))
	if troopType == "" || err != nil || e != nil || quantity < 1 {
		fmt.Fprintf(w, "invalid selction")
		return
	}

	create := models.CreateTroop(username, troopType, quantity, level)
	if create == false {
		fmt.Fprintf(w, "could not create trop")
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
		fmt.Fprintf(w, "bad req")
		return
	}

	defenseType := r.FormValue("defense_type")
	am, err := strconv.Atoi(r.FormValue("amount"))
	if defenseType == "" || err != nil || am < 1 {
		fmt.Fprintf(w, "invalid selection")
		return
	}

	create := models.CreateDefense(username, defenseType, am)

	if create == false {
		fmt.Fprintf(w, "could not create trop")
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
		fmt.Fprintf(w, "bad req")
		return
	}

	buildingType := r.FormValue("building_type")
	x, err := strconv.Atoi(r.FormValue("x"))
	y, e := strconv.Atoi(r.FormValue("y"))
	if buildingType == "" || err != nil || e != nil {
		fmt.Fprintf(w, "invalid selection")
		return
	}
	create := models.CreateBuilding(username, buildingType, x, y)

	if create == false {
		fmt.Fprintf(w, "could not create trop")
	}

	w.Header().Set("HX-Redirect", "/web/app/shop")
	w.WriteHeader(http.StatusOK)
}
