package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
