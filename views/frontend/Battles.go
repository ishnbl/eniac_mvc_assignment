package frontend

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
)

type BatData struct {
	Opps []models.RetScoutVillage
}

func BattlesHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	data := BatData{
		Opps: models.GetAllVillages(username),
	}

	tmpl := template.Must(template.ParseGlob("views/templates/*.html"))
	err = tmpl.ExecuteTemplate(w, "battles.html", data)
	if err != nil {
		fmt.Fprintf(w, "could not render template")
	}
}

type RenderBattle struct {
	DefVill models.Village
	AtVill  models.Village
	DefDef  []models.RetDefenses
	AtTroop []models.RetTroops
}

func DoBattleHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	db := models.DB
	var village models.Village
	villId, err := strconv.Atoi(mux.Vars(r)["villageID"])

	if err != nil {
		return
	}
	db.Where(&models.Village{ID: uint(villId)}).First(&village)
	att := models.GetVillage(username)
	troops := models.GetTroops(username)
	var user models.User
	db.Where(&models.User{ID: village.UserID}).First(&user)
	defenses := models.GetDefenses(user.Username)

	data := RenderBattle{
		DefVill: village,
		AtVill:  att,
		DefDef:  defenses,
		AtTroop: troops,
	}
	tmpl := template.Must(template.ParseGlob("views/templates/*.html"))
	err = tmpl.ExecuteTemplate(w, "battle.html", data)
	if err != nil {
		fmt.Fprintf(w, "could not render template")
	}

}

type DoFight struct {
	DefenderId uint
	Config     []models.FightConf
}

func FightHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	username, err := controller.GetUsernameFromToken(cookie.Value)
	if err != nil {
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "bad req")
		return
	}

	defenderID, err := strconv.Atoi(r.FormValue("defender_id"))
	defenderID_u := uint(defenderID)
	if err != nil {
		fmt.Fprintf(w, "bad req")
		return
	}

	amounts := r.Form["amount"]
	troopTypes := r.Form["troop_type"]
	troopLevels := r.Form["troop_level"]
	skills := r.Form["skill"]

	var config []models.FightConf
	for i := 0; i < len(amounts); i++ {
		amount, err := strconv.Atoi(amounts[i])
		if err != nil {
			continue
		}
		level, err := strconv.Atoi(troopLevels[i])
		if err != nil {
			continue
		}
		config = append(config, models.FightConf{
			TypeTroop:  troopTypes[i],
			LevelTroop: level,
			AmountDep:  amount,
			CapType:    skills[i],
		})
	}

	won, loot := models.CreateFight(username, defenderID_u, config)
	if won {
		fmt.Fprintf(w, "Victory! You looted %d resources.", loot)
	} else {
		fmt.Fprintf(w, "Defeat")
	}
}
