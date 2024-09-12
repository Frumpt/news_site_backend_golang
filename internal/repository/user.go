package repository

import (
	"NewsBack/internal/db"
	"NewsBack/internal/domain"
	"NewsBack/sqlc/database"
	"context"
	"fmt"
)

type userDateBase struct {
	DB *db.DB
}

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindOne(int) (domain.User, error)
	Save(domain.User) (domain.User, error)
	DeleteById(int) (domain.User, error)
}

func NewUserRepository(DB *db.DB) UserRepository {
	return &userDateBase{DB: DB}
}

func (udb *userDateBase) FindAll() ([]domain.User, error) {
	var users []domain.User
	ctx := context.Background()
	users, err := udb.DB.UseCase.GetUsers(ctx)

	return users, err
}

func (udb *userDateBase) FindOne(id int) (domain.User, error) {
	var users domain.User
	ctx := context.Background()
	users, err := udb.DB.UseCase.GetUser(ctx, id)

	return users, err
}

func (udb *userDateBase) Save(user domain.User) (domain.User, error) {
	var users domain.User
	ctx := context.Background()
	users, err := udb.DB.UseCase.CreateUser(ctx, database.CreateUserParams{ID: user.ID, Name: user.Name, Password: *user.Password, UserRoleID: user.UserRoleID})

	return users, err
}

func (udb *userDateBase) DeleteById(id int) (domain.User, error) {
	var users domain.User
	ctx := context.Background()
	rows, err := udb.DB.UseCase.DeleteUser(ctx, id)
	fmt.Printf("%v, %v, %v", rows, err, "deleted")
	if rows == 0 {
		return users, fmt.Errorf("record not found")
	}
	return users, err
}
