package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init_DB() {
	db_Str := os.Getenv("DB_STRING")
	db, err := gorm.Open(postgres.Open(db_Str), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
