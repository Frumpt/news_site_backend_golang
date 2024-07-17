package main

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func dropTable() {
	db.Connect(configDB)
	db := db.DataBase
	errMigrate := db.Migrator().DropTable(&models.Roles{}, &models.Users{}, &models.News{}, &models.Comments{}, &models.Tags{}, &models.NewsTags{})

	if errMigrate != nil {
		panic(errMigrate)
	} else {
		println("Table dropped")
	}

}
