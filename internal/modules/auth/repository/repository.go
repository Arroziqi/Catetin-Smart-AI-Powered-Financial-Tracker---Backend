package repository

import "catetin-backend/internal/modules/auth/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}