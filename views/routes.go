package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
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
		UserID:           user.ID,
		VillageLevel:     1,
		Gold:             20,
		Oil:              50,
		Money:            5000,
		FarmLand:         20,
		Mines:            0,
		LevelConstraints: datatypes.JSON(`{}`),
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

type returnBuilding struct {
	X      int
	Y      int
	Type   string
	Level  int
	Width  int
	Height int
}

type villageResp struct {
	Village   models.Village
	Buildings []returnBuilding
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
	buildingsId := []models.BuildingsVillageMapping{}
	db.Where(&models.BuildingsVillageMapping{VillageID: village.ID}).Find(&buildingsId)
	buildings := []returnBuilding{}

	for i := 0; i < len(buildingsId); i++ {
		var building models.Buildings
		db.Where(&models.Buildings{ID: buildingsId[i].BuildingsID}).First(&building)
		buildings = append(buildings, returnBuilding{
			X:      buildingsId[i].X,
			Y:      buildingsId[i].Y,
			Type:   building.Type,
			Width:  building.Width,
			Height: building.Height,
		})

	}

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

	var buildingType models.Buildings
	db.Where(&models.Buildings{Type: payload.Type}).First(&buildingType)
	if buildingType.ID == 0 {
		fmt.Fprintf(w, "Invalid building type")
		return
	}
	if buildingType.Cost > village.Money {
		fmt.Fprintf(w, "Not enough money")
		return
	}
	village.Money -= buildingType.Cost
	building := models.BuildingsVillageMapping{
		VillageID:   village.ID,
		BuildingsID: buildingType.ID,
		X:           payload.X,
		Y:           payload.Y,
	}

	buildings := []models.BuildingsVillageMapping{}
	db.Where(&models.BuildingsVillageMapping{VillageID: village.ID}).Find(&buildings)

	for _, b := range buildings {
		if (payload.X < b.X+buildingType.Width) && (b.X < payload.X+buildingType.Width) &&
			(payload.Y < b.Y+buildingType.Height) && (b.Y < payload.Y+buildingType.Height) {
			fmt.Fprintf(w, "Building overlaps with existing building")
			return
		}
	}

	db.Save(&village)
	result := db.Create(&building)
	if result.Error != nil {
		fmt.Fprintf(w, "Error creating building")
		return
	}
	fmt.Fprintf(w, "Building created successfully")
}

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

	var defenseType models.Defenses

	db.Where(&models.Defenses{Type: payload.Type}).First(&defenseType)

	if defenseType.ID == 0 {
		fmt.Fprintf(w, "invalid type")
		return
	}
	defense := models.VillageDefMapping{
		VillageID:  village.ID,
		Amount:     payload.Amount,
		DefensesID: defenseType.ID,
	}

	if defenseType.Cost > village.Money {
		fmt.Fprintf(w, "Erroro ")
		return
	}

	village.Money = village.Money - defenseType.Cost
	result := db.Create(&defense)
	if result.Error != nil {
		fmt.Fprintf(w, "Error creating defense")
		return
	}
	fmt.Fprintf(w, "CreateDefenses endpoint")
	db.Save(&village)
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

	var troopType models.Troops

	db.Where(&models.Troops{Type: payload.Type}).First(&troopType)
	if troopType.ID == 0 {
		fmt.Fprintf(w, "invalid type")
		return
	}
	if troopType.Cost > village.Money {
		fmt.Fprintf(w, "money isues")
		return
	}

	saveTroop := models.TroopVillageMapping{
		VillageID: village.ID,
		TroopsID:  troopType.ID,
		Quantity:  payload.Quantity,
	}

	db.Create(&saveTroop)
	village.Money = village.Money - troopType.Cost
	db.Save(&village)
}
