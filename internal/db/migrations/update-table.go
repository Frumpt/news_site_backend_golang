package main

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func updateTable() {
	db.Connect(configDB)
	db := db.DataBase
	errMigrate := db.Migrator().AlterColumn(&models.News{}, "tag_id")

	if errMigrate != nil {
		panic(errMigrate)
	} else {
		println("Table update")
	}
}
