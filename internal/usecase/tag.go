package usecase

import (
	"NewsBack/internal/api/Router"
	"NewsBack/internal/domain"
	"NewsBack/internal/repository"
)

type TagUseCase struct {
	TagRepo repository.TagRepository
}

func NewTagUseCase(repo repository.TagRepository) Router.TagUseCase {
	return &TagUseCase{TagRepo: repo}
}

func (uc *TagUseCase) FindAll() ([]domain.Tag, error) {
	Tags, err := uc.TagRepo.FindAll()
	return Tags, err
}

func (uc *TagUseCase) FindOne(id int) (domain.Tag, error) {
	Tag, err := uc.TagRepo.FindOne(id)
	return Tag, err
}

func (uc *TagUseCase) Save(Tag domain.Tag) (domain.Tag, error) {
	Tag, err := uc.TagRepo.Save(Tag)
	return Tag, err
}

func (uc *TagUseCase) DeleteById(id int) (domain.Tag, error) {
	Tag, err := uc.TagRepo.DeleteById(id)
	return Tag, err
}
