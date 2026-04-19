package usecase

import (
	"catetin-backend/internal/modules/transaction/domain"
	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/repository"
)

type GetTransactionUsecase interface {
	Execute(userID uint, query dto.GetTransactionQuery) ([]domain.Transaction, int64, error)
}

type getTransactionUsecase struct {
	repo repository.ITransactionRepository
}

func NewGetTransactionUsecase(repo repository.ITransactionRepository) GetTransactionUsecase {
	return &getTransactionUsecase{repo: repo}
}

func (u *getTransactionUsecase) Execute(userID uint, query dto.GetTransactionQuery) ([]domain.Transaction, int64, error) {
	return u.repo.FindByUser(userID, query)
}
