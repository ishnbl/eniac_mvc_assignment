package views

import (
	"encoding/json"
	"fmt"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"gorm.io/datatypes"
	"net/http"
)

type RegisterResp struct {
	Username string
	Name     string
	Password string
}

type LoginResp struct {
	Username string
	Password string
}


type LoginRet struct {
	Token string
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterResp
	db := models.DB
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(w, "Invalid Payload")
	}
	hashedPassword, err := Hash(payload.Password)
	if err != nil {
		fmt.Fprintf(w, "Error hashing password")
	}
	fmt.Println("Hashed Password:", hashedPassword)
	user := models.User{Name: payload.Name, Username: payload.Username, HashedPassword: hashedPassword}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Fprintf(w, "error creating user")
		return
	}

	village := models.Village{
		UserID:       user.ID,
		VillageLevel: 1,
		Gold:         20,
		Oil:          50,
		Money:        5000,
		FarmLand:     20,
		Mines:        0,
		Map:          datatypes.JSON(`{}`),
	}
	result = db.Create(&village)
	if result.Error != nil {
		fmt.Fprintf(w, "error creating village")
		return
	}

	fmt.Fprintf(w, "Register endpoint")

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginHandler called")
	var payload LoginResp
	db := models.DB
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(w, "Invalid Payload")
	}
	var user models.User
	db.Where(&models.User{Username: payload.Username}).First(&user)
	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(payload.Password))
	if passwordMatch != nil {
		fmt.Fprintf(w, "Login failed")
	}
	jwtToken, err := controller.CreateToken(payload.Username)
	if err != nil {
		fmt.Fprintf(w, "Error creating token")
	}
	fmt.Println("Password match:", passwordMatch)
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
	jsonResp, err := json.Marshal(LoginRet{Token: jwtToken})
	if err != nil {
		fmt.Fprintf(w, "Error creating JSON response")
	}
	w.Write(jsonResp)
}

type villageResp struct {
	Village   models.Village
	Buildings []models.Buildings
}


func VillageHandler(w http.ResponseWriter, r *http.Request) {
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}
	db := models.DB
	var user models.User
	var village models.Village

	db.Where(&models.User{Username: username}).First(&user)
	db.Where(&models.Village{UserID: user.ID}).First(&village)
	buildings := []models.Buildings{}
	db.Where(&models.Buildings{VillageID: village.ID}).Find(&buildings)
	jsonResp, err := json.Marshal(villageResp{Village: village, Buildings: buildings})
	if err != nil {
		fmt.Fprintf(w, "Error creating JSON response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}


type CreateBuildingReq struct {
	BuildingLevel int
	X             int
	Y             int
	Width         int
	Height        int
}

func CreateBuildingsHandler(w http.ResponseWriter, r *http.Request) {
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}
	db := models.DB
	var user models.User
	var village models.Village

	db.Where(&models.User{Username: username}).First(&user)
	db.Where(&models.Village{UserID: user.ID}).First(&village)
	var payload CreateBuildingReq
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(w, "Invalid Payload")
		return
	}
	building := models.Buildings{
		VillageID: village.ID,
		BuildingLevel: payload.BuildingLevel,
		X: payload.X,
		Y: payload.Y,
		Width: payload.Width,
		Height: payload.Height,
	}

	buildings := []models.Buildings{}
	db.Where(&models.Buildings{VillageID: village.ID}).Find(&buildings)

	for _, b := range buildings {
		if (payload.X < b.X + b.Width) && (b.X < payload.X + payload.Width) &&
			(payload.Y < b.Y + b.Height) && (b.Y < payload.Y + payload.Height) {
			fmt.Fprintf(w, "Building overlaps with existing building")
			return
		}
	}
	result := db.Create(&building)
	if result.Error != nil {
		fmt.Fprintf(w, "Error creating building")
		return
	}
	fmt.Fprintf(w, "Building created successfully")
}


type CreateDefensesReq struct {
	Type           string
	DefensivePower int
	Amount         int
}

func CreateDefensesHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateDefensesReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprintf(w, "Invalid Payload")
		return
	}
	db := models.DB
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}
	var user models.User
	var village models.Village

	db.Where(&models.User{Username: username}).First(&user)
	db.Where(&models.Village{UserID: user.ID}).First(&village)

	defense := models.Defenses{
		VillageID: village.ID,
		Type: payload.Type,
		DefensivePower: payload.DefensivePower,
		Amount: payload.Amount,
	}
	result := db.Create(&defense)
	if result.Error != nil {
		fmt.Fprintf(w, "Error creating defense")
		return
	}
	fmt.Fprintf(w, "CreateDefenses endpoint")
}

type TroopConstraints struct {
	Type           string
	Health         int
	OffensivePower int
	Level          int
	Quantity       int
	Cost           int
}
//abhi ke liye dummy constraints, afterr db reveiw will fetch from table
var Archer = TroopConstraints{
	Type: "Archer",
	Health: 100,
	OffensivePower: 50,
	Level: 1,
	Cost: 10,
}

var Assasin = TroopConstraints{
	Type: "Asassin",
	Health: 150,
	OffensivePower: 70,
	Level: 1,
	Quantity: 5,
	Cost: 20,
}

var Cavalry = TroopConstraints{
	Type: "Cavalry",
	Health: 200,
	OffensivePower: 100,
	Level: 1,
	Quantity: 3,
	Cost: 30,
}

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
	db := models.DB
	username, err := controller.GetUsernameFromToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "Error getting username from token")
		return
	}
	var user models.User
	var village models.Village

	db.Where(&models.User{Username: username}).First(&user)
	db.Where(&models.Village{UserID: user.ID}).First(&village)

	troop := models.Troops{}
	if payload.Type == "Archer" {
	troop = models.Troops{
		VillageID: village.ID,
		Type: payload.Type,
		Health: Archer.Health,
		OffensivePower: Archer.OffensivePower,
		Level: Archer.Level,
		Quantity: payload.Quantity,
	}
	}else if payload.Type == "Assasin" {
	troop = models.Troops{
		VillageID: village.ID,
		Type: payload.Type,
		Health: Assasin.Health,
		OffensivePower: Assasin.OffensivePower,
		Level: Assasin.Level,
		Quantity: payload.Quantity,
	}}else if payload.Type == "Cavalry" {
	troop = models.Troops{
		VillageID: village.ID,
		Type: Cavalry.Type,
		Health: Cavalry.Health,
		OffensivePower: Cavalry.OffensivePower,
		Level: Cavalry.Level,
		Quantity: payload.Quantity,
	}
}

	if troop.Type == "Archer" {
		if troop.Quantity * Archer.Cost > village.Money {
			fmt.Fprintf(w, "Not enough money to create troops")
			return
		}
		village.Money -= troop.Quantity * Archer.Cost
	}else if troop.Type == "Assasin" {
		if troop.Quantity * Assasin.Cost > village.Money {
			fmt.Fprintf(w, "Not enough money to create troops")
			return
		}
		village.Money -= troop.Quantity * Assasin.Cost
	}else if troop.Type == "Cavalry" {
		if troop.Quantity * Cavalry.Cost > village.Money {
			fmt.Fprintf(w, "Not enough money to create troops")
			return
		}
		village.Money -= troop.Quantity * Cavalry.Cost
	}

	db.Save(&village)
	result := db.Create(&troop)
	if result.Error != nil {
		fmt.Fprintf(w, "Error creating troop")
		return
	}
	fmt.Fprintf(w, "CreateTroops endpoint")
}