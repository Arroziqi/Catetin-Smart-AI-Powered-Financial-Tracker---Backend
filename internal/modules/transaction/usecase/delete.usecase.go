package usecase

import (
	"catetin-backend/internal/modules/transaction/repository"
)

type DeleteTransactionUsecase interface {
	Execute(id uint, userID uint) error
}

type deleteTransactionUsecase struct {
	repo repository.ITransactionRepository
}

func NewDeleteTransactionUsecase(repo repository.ITransactionRepository) DeleteTransactionUsecase {
	return &deleteTransactionUsecase{repo: repo}
}

func (u *deleteTransactionUsecase) Execute(id uint, userID uint) error {
	return u.repo.Delete(id, userID)
}
