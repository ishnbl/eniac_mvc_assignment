package frontend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/ishnbl/eniac_mvc_assignment/controller"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("views/templates/*.html"))
	err := tmpl.ExecuteTemplate(w, "auth.html", nil)
	if err != nil {
		fmt.Fprintf(w, "could not render template")
	}
}
func LoginViewHtmx(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginHandler called")

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "bad req")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	db := models.DB
	var user models.User
	result := db.Where(&models.User{Username: username}).First(&user)
	if result.Error != nil {
		fmt.Fprintf(w, "account does not exist")
		return
	}

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if passwordMatch != nil {
		fmt.Fprintf(w, "unauthorized")
		return
	}

	jwtToken, err := controller.CreateToken(username)
	if err != nil {
		fmt.Fprintf(w, "error try again")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("HX-Redirect", "/web/app/village")
	w.WriteHeader(http.StatusOK)
}

func RegisterHandlerHtmx(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "bad req")
		return
	}

	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" || name == "" {
		fmt.Fprintf(w, "empty")
		return
	}

	db := models.DB

	hashedPassword, err := controller.Hash(password)
	if err != nil {
		fmt.Fprintf(w, "error")
		return
	}

	user := models.User{Name: name, Username: username, HashedPassword: hashedPassword}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Fprintf(w, "could not create user")
		return
	}

	constraintsJSON, _ := json.Marshal(models.LCons[0])
	village := models.Village{
		UserID:           user.ID,
		VillageLevel:     1,
		Gold:             20,
		Oil:              50,
		Money:            5000,
		FarmLand:         20,
		Mines:            0,
		LevelConstraints: datatypes.JSON(constraintsJSON),
	}

	result = db.Create(&village)
	if result.Error != nil {
		fmt.Fprintf(w, "no village for you")
		return
	}

	w.Header().Set("HX-Redirect", "/web/login")
	w.WriteHeader(http.StatusOK)
}
