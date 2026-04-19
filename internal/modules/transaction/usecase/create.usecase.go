package usecase

import (
	"catetin-backend/internal/modules/transaction/domain"
	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/repository"
)

type CreateTransactionUsecase interface {
	Execute(userID uint, req dto.CreateTransactionRequest) error
}

type createTransactionUsecase struct {
	repo repository.ITransactionRepository
}

func NewCreateTransactionUsecase(repo repository.ITransactionRepository) CreateTransactionUsecase {
	return &createTransactionUsecase{repo: repo}
}

func (u *createTransactionUsecase) Execute(userID uint, req dto.CreateTransactionRequest) error {
	transaction := domain.Transaction{
		UserID:   userID,
		Type:     req.Type,
		Amount:   req.Amount,
		Category: req.Category,
		Note:     req.Note,
		Date:     req.Date,
	}

	return u.repo.Create(transaction)
}
