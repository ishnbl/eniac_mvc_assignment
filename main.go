package main

import (
	"fmt"
	"github.com/ishnbl/eniac_mvc_assignment/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, MVC!")
	db_Str := os.Getenv("DB_STRING")
	db, err := gorm.Open(postgres.Open(db_Str), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.User{}, &models.Village{}, &models.Defenses{}, &models.Fights{}, &models.Troops{}, &models.Fights{}, &models.Buildings{})

}
