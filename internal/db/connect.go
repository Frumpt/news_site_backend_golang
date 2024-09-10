package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func Connect(config string) {

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DataBase = db
}
