package usecase

import (
	"NewsBack/internal/api/Router"
	"NewsBack/internal/domain"
	"NewsBack/internal/repository"
)

type CommentUseCase struct {
	CommentRepo repository.CommentRepository
}

func NewCommentUseCase(repo repository.CommentRepository) Router.CommentUseCase {
	return &CommentUseCase{CommentRepo: repo}
}

func (uc *CommentUseCase) FindAll() ([]domain.Comment, error) {
	Comments, err := uc.CommentRepo.FindAll()
	return Comments, err
}

func (uc *CommentUseCase) FindOne(id int) (domain.Comment, error) {
	Comment, err := uc.CommentRepo.FindOne(id)
	return Comment, err
}

func (uc *CommentUseCase) Save(Comment domain.Comment) (domain.Comment, error) {
	Comment, err := uc.CommentRepo.Save(Comment)
	return Comment, err
}

func (uc *CommentUseCase) DeleteById(id int) (domain.Comment, error) {
	Comment, err := uc.CommentRepo.DeleteById(id)
	return Comment, err
}
