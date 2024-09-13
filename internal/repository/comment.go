package repository

import (
	"NewsBack/internal/db"
	"NewsBack/internal/domain"
	"NewsBack/sqlc/database"
	"context"
	"fmt"
)

type CommentDateBase struct {
	DB *db.DB
}

type CommentRepository interface {
	FindAll() ([]domain.Comment, error)
	FindOne(int) (domain.Comment, error)
	Save(domain.Comment) (domain.Comment, error)
	DeleteById(int) (domain.Comment, error)
}

func NewCommentRepository(DB *db.DB) CommentRepository {
	return &CommentDateBase{DB: DB}
}

func (udb *CommentDateBase) FindAll() ([]domain.Comment, error) {
	var Comments []domain.Comment
	ctx := context.Background()
	Comments, err := udb.DB.UseCase.GetComments(ctx)

	return Comments, err
}

func (udb *CommentDateBase) FindOne(id int) (domain.Comment, error) {
	var Comments domain.Comment
	ctx := context.Background()
	Comments, err := udb.DB.UseCase.GetComment(ctx, id)

	return Comments, err
}

func (udb *CommentDateBase) Save(Comment domain.Comment) (domain.Comment, error) {
	var Comments domain.Comment
	ctx := context.Background()
	Comments, err := udb.DB.UseCase.CreateComment(ctx, database.CreateCommentParams{ID: Comment.ID, Name: Comment.Name, NewsID: Comment.NewsID, UserID: Comment.UserID, Description: Comment.Description})

	return Comments, err
}

func (udb *CommentDateBase) DeleteById(id int) (domain.Comment, error) {
	var Comments domain.Comment
	ctx := context.Background()
	rows, err := udb.DB.UseCase.DeleteComment(ctx, id)
	if rows == 0 {
		return Comments, fmt.Errorf("record not found")
	}
	return Comments, err
}
