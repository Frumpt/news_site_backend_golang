package repository

import (
	"NewsBack/internal/db"
	"NewsBack/internal/domain"
	"NewsBack/sqlc/database"
	"context"
	"fmt"
)

type NewsDateBase struct {
	DB *db.DB
}

type NewsRepository interface {
	FindAll() ([]domain.News, error)
	FindOne(int) (domain.News, error)
	Save(domain.News) (domain.News, error)
	DeleteById(int) (domain.News, error)
}

func NewNewsRepository(DB *db.DB) NewsRepository {
	return &NewsDateBase{DB: DB}
}

func (udb *NewsDateBase) FindAll() ([]domain.News, error) {
	var News []domain.News
	ctx := context.Background()
	News, err := udb.DB.UseCase.GetNews(ctx)

	return News, err
}

func (udb *NewsDateBase) FindOne(id int) (domain.News, error) {
	var News domain.News
	ctx := context.Background()
	News, err := udb.DB.UseCase.GetNew(ctx, id)

	return News, err
}

func (udb *NewsDateBase) Save(New domain.News) (domain.News, error) {
	var News domain.News
	ctx := context.Background()
	News, err := udb.DB.UseCase.CreateNew(ctx, database.CreateNewParams{ID: New.ID, Title: New.Title, UserID: New.UserID, Description: New.Description, NameImage: New.NameImage})

	return News, err
}

func (udb *NewsDateBase) DeleteById(id int) (domain.News, error) {
	var News domain.News
	ctx := context.Background()
	rows, err := udb.DB.UseCase.DeleteNews(ctx, id)
	if rows == 0 {
		return News, fmt.Errorf("record not found")
	}
	return News, err
}
