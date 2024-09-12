package db

import (
	"NewsBack/sqlc/database"
	"context"
	"github.com/jackc/pgx/v5"
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

type DB struct {
	UseCase *database.Queries
}

func ConnectPGX(config string) (*DB, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config)
	if err != nil {
		return nil, err
	}

	//defer conn.Close(ctx)

	Queries := database.New(conn)

	db := DB{Queries}

	return &db, nil
}
