package repository

import (
	"NewsBack/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

type CommentDateBase struct {
	DB *gorm.DB
}

type CommentRepository interface {
	FindAll() ([]domain.Comment, error)
	FindOne(int) (domain.Comment, error)
	Save(domain.Comment) (domain.Comment, error)
	DeleteById(int) (domain.Comment, error)
}

func NewCommentRepository(DB *gorm.DB) CommentRepository {
	return &CommentDateBase{DB: DB}
}

func (udb *CommentDateBase) FindAll() ([]domain.Comment, error) {
	var Comments []domain.Comment
	err := udb.DB.Select("ID", "UserID", "Name", "Description", "NewsID").Find(&Comments).Error

	return Comments, err
}

func (udb *CommentDateBase) FindOne(id int) (domain.Comment, error) {
	var Comment domain.Comment
	err := udb.DB.First(&Comment, id).Error

	return Comment, err
}

func (udb *CommentDateBase) Save(Comment domain.Comment) (domain.Comment, error) {
	err := udb.DB.Create(&Comment).Error

	return Comment, err
}

func (udb *CommentDateBase) DeleteById(id int) (domain.Comment, error) {
	var Comment domain.Comment
	var err error
	rows := udb.DB.Where("id = ?", id).Delete(&Comment).RowsAffected
	if rows == 0 {
		err = fmt.Errorf("not found")
	}
	return Comment, err

}
