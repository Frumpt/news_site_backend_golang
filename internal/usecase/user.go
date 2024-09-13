package usecase

import (
	"NewsBack/internal/api/Router"
	"NewsBack/internal/domain"
	"NewsBack/internal/repository"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) Router.UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (uc *UserUseCase) FindAll() ([]domain.User, error) {
	users, err := uc.userRepo.FindAll()
	return users, err
}

func (uc *UserUseCase) FindOne(id int) (domain.User, error) {
	user, err := uc.userRepo.FindOne(id)
	return user, err
}

func (uc *UserUseCase) Save(user domain.User) (domain.User, error) {
	user, err := uc.userRepo.Save(user)
	return user, err
}

func (uc *UserUseCase) DeleteById(id int) (domain.User, error) {
	user, err := uc.userRepo.DeleteById(id)
	return user, err
}
