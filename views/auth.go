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
