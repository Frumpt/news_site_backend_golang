package repository

import (
	"NewsBack/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.0 --name=DB
type DB interface {
	NewDB() *gorm.DB
}

type userDateBase struct {
	DB *gorm.DB
}

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindOne(int) (domain.User, error)
	Save(domain.User) (domain.User, error)
	DeleteById(int) (domain.User, error)
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userDateBase{DB: DB}
}

func (udb *userDateBase) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := udb.DB.Select("Id", "Name", "user_role_id").Find(&users).Error

	return users, err
}

func (udb *userDateBase) FindOne(id int) (domain.User, error) {
	var user domain.User
	err := udb.DB.First(&user, id).Error

	return user, err
}

func (udb *userDateBase) Save(user domain.User) (domain.User, error) {
	err := udb.DB.Create(&user).Error

	return user, err
}

func (udb *userDateBase) DeleteById(id int) (domain.User, error) {
	var user domain.User
	var err error
	rows := udb.DB.Where("id = ?", id).Delete(&user).RowsAffected
	if rows == 0 {
		err = fmt.Errorf("not found")
	}
	return user, err

}
