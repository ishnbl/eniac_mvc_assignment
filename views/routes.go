package views

import (
	"encoding/json"
	"fmt"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"golang.org/x/crypto/bcrypt"
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
	result := db.Create(&models.User{Name: payload.Name, Username: payload.Username, HashedPassword: hashedPassword})
	if result.Error != nil {
		fmt.Fprintf(w, "error creating")
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

	fmt.Println("Password match:", passwordMatch)
	fmt.Fprintf(w, "Login endpoint")
}
