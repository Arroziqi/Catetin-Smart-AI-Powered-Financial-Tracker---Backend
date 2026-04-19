package usecase

import (
	"catetin-backend/internal/modules/transaction/domain"
	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/repository"
)

type UpdateTransactionUsecase interface {
	Execute(id uint, userID uint, req dto.UpdateTransactionRequest) error
}

type updateTransactionUsecase struct {
	repo repository.ITransactionRepository
}

func NewUpdateTransactionUsecase(repo repository.ITransactionRepository) UpdateTransactionUsecase {
	return &updateTransactionUsecase{repo: repo}
}

func (u *updateTransactionUsecase) Execute(id uint, userID uint, req dto.UpdateTransactionRequest) error {
	transaction := domain.Transaction{
		Type:     req.Type,
		Amount:   req.Amount,
		Category: req.Category,
		Note:     req.Note,
		Date:     req.Date,
	}

	return u.repo.Update(id, userID, transaction)
}
