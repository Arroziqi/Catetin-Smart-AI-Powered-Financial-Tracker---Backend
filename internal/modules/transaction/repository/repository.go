package repository

import (
	"catetin-backend/internal/modules/transaction/domain"
	"catetin-backend/internal/modules/transaction/dto"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(data domain.Transaction) error {
	return r.db.Create(&data).Error
}

func (r *repository) FindByUser(userID uint, query dto.GetTransactionQuery) ([]domain.Transaction, int64, error) {
	var transactions []domain.Transaction
	var total int64

	db := r.db.Model(&domain.Transaction{}).Where("user_id = ?", userID)

	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}

	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}

	if query.StartDate != "" && query.EndDate != "" {
		db = db.Where("date BETWEEN ? AND ?", query.StartDate, query.EndDate)
	} else if query.StartDate != "" {
		db = db.Where("date >= ?", query.StartDate)
	} else if query.EndDate != "" {
		db = db.Where("date <= ?", query.EndDate)
	}

	// Count total records
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	limit := query.Limit
	if limit <= 0 {
		limit = 10
	}
	page := query.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	db = db.Order("date DESC").Limit(limit).Offset(offset)

	if err := db.Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

func (r *repository) Update(id uint, userID uint, data domain.Transaction) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Updates(&data).Error
}

func (r *repository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Transaction{}).Error
}
