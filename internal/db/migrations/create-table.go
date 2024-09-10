package main

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"NewsBack/internal/service/roles"
	"encoding/json"
)

func createTable() {
	db.Connect(configDB)
	db := db.DataBase
	errMigrate := db.Migrator().CreateTable(&models.Roles{}, &models.Users{}, &models.News{}, &models.Comments{}, &models.Tags{}, &models.NewsTags{})

	if errMigrate != nil {
		panic(errMigrate)
	} else {

		bytes, err := json.Marshal(struct {
			name string
			id   int
		}{"User", 1})
		roles.CreateRole(bytes)
		bytes, err = json.Marshal(struct {
			name string
			id   int
		}{"Admin", 1})
		roles.CreateRole(bytes)
		if err != nil {
			panic(err)
		}

		println("Table created")
	}
}
