package repository

import (
	"catetin-backend/internal/modules/transaction/domain"
	"catetin-backend/internal/modules/transaction/dto"
)

type ITransactionRepository interface {
	Create(data domain.Transaction) error
	FindByUser(userID uint, query dto.GetTransactionQuery) ([]domain.Transaction, int64, error)
	Update(id uint, userID uint, data domain.Transaction) error
	Delete(id uint, userID uint) error
}
