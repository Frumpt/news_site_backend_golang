package usecase

import (
	"NewsBack/internal/api/Router"
	"NewsBack/internal/domain"
	"NewsBack/internal/repository"
)

type NewsUseCase struct {
	NewsRepo repository.NewsRepository
}

func NewNewsUseCase(repo repository.NewsRepository) Router.NewsUseCase {
	return &NewsUseCase{NewsRepo: repo}
}

func (uc *NewsUseCase) FindAll() ([]domain.News, error) {
	News, err := uc.NewsRepo.FindAll()
	return News, err
}

func (uc *NewsUseCase) FindOne(id int) (domain.News, error) {
	News, err := uc.NewsRepo.FindOne(id)
	return News, err
}

func (uc *NewsUseCase) Save(News domain.News) (domain.News, error) {
	News, err := uc.NewsRepo.Save(News)
	return News, err
}

func (uc *NewsUseCase) DeleteById(id int) (domain.News, error) {
	News, err := uc.NewsRepo.DeleteById(id)
	return News, err
}
