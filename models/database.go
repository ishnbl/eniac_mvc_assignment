package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB() {
	db_Str := os.Getenv("DB_STRING")
	if db_Str == "" {
		db_Str = "host=localhost user=admin password=secret dbname=mvc port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(db_Str), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
