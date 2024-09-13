package repository

import (
	"NewsBack/internal/db"
	"NewsBack/internal/domain"
	"NewsBack/sqlc/database"
	"context"
	"fmt"
)

type TagDateBase struct {
	DB *db.DB
}

type TagRepository interface {
	FindAll() ([]domain.Tag, error)
	FindOne(int) (domain.Tag, error)
	Save(domain.Tag) (domain.Tag, error)
	DeleteById(int) (domain.Tag, error)
}

func NewTagRepository(DB *db.DB) TagRepository {
	return &TagDateBase{DB: DB}
}

func (udb *TagDateBase) FindAll() ([]domain.Tag, error) {
	var Tags []domain.Tag
	ctx := context.Background()
	Tags, err := udb.DB.UseCase.GetTags(ctx)

	return Tags, err
}

func (udb *TagDateBase) FindOne(id int) (domain.Tag, error) {
	var Tags domain.Tag
	ctx := context.Background()
	Tags, err := udb.DB.UseCase.GetTag(ctx, id)

	return Tags, err
}

func (udb *TagDateBase) Save(Tag domain.Tag) (domain.Tag, error) {
	var Tags domain.Tag
	ctx := context.Background()
	Tags, err := udb.DB.UseCase.CreateTag(ctx, database.CreateTagParams{ID: Tag.ID, Name: Tag.Name})

	return Tags, err
}

func (udb *TagDateBase) DeleteById(id int) (domain.Tag, error) {
	var Tags domain.Tag
	ctx := context.Background()
	rows, err := udb.DB.UseCase.DeleteTags(ctx, id)
	if rows == 0 {
		return Tags, fmt.Errorf("record not found")
	}
	return Tags, err
}
