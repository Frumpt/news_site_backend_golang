package repository

import (
	"NewsBack/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

type TagDateBase struct {
	DB *gorm.DB
}

type TagRepository interface {
	FindAll() ([]domain.Tag, error)
	FindOne(int) (domain.Tag, error)
	Save(domain.Tag) (domain.Tag, error)
	DeleteById(int) (domain.Tag, error)
}

func NewTagRepository(DB *gorm.DB) TagRepository {
	return &TagDateBase{DB: DB}
}

func (udb *TagDateBase) FindAll() ([]domain.Tag, error) {
	var Tags []domain.Tag
	err := udb.DB.Select("Id", "Name").Find(&Tags).Error

	return Tags, err
}

func (udb *TagDateBase) FindOne(id int) (domain.Tag, error) {
	var Tag domain.Tag
	err := udb.DB.First(&Tag, id).Error

	return Tag, err
}

func (udb *TagDateBase) Save(Tag domain.Tag) (domain.Tag, error) {
	err := udb.DB.Create(&Tag).Error

	return Tag, err
}

func (udb *TagDateBase) DeleteById(id int) (domain.Tag, error) {
	var Tag domain.Tag
	var err error
	rows := udb.DB.Where("id = ?", id).Delete(&Tag).RowsAffected
	if rows == 0 {
		err = fmt.Errorf("not found")
	}
	return Tag, err

}
