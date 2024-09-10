package repository

import (
	"NewsBack/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

type NewsDateBase struct {
	DB *gorm.DB
}

type NewsRepository interface {
	FindAll() ([]domain.News, error)
	FindOne(int) (domain.News, error)
	Save(domain.News) (domain.News, error)
	DeleteById(int) (domain.News, error)
}

func NewNewsRepository(DB *gorm.DB) NewsRepository {
	return &NewsDateBase{DB: DB}
}

func (udb *NewsDateBase) FindAll() ([]domain.News, error) {
	var News []domain.News
	err := udb.DB.Select("ID", "UserID", "Title", "Description", "NameImage").Find(&News).Error

	return News, err
}

func (udb *NewsDateBase) FindOne(id int) (domain.News, error) {
	var News domain.News
	err := udb.DB.First(&News, id).Error

	return News, err
}

func (udb *NewsDateBase) Save(News domain.News) (domain.News, error) {
	err := udb.DB.Create(&News).Error

	return News, err
}

func (udb *NewsDateBase) DeleteById(id int) (domain.News, error) {
	var News domain.News
	var err error
	rows := udb.DB.Where("id = ?", id).Delete(&News).RowsAffected
	if rows == 0 {
		err = fmt.Errorf("not found")
	}
	return News, err

}
