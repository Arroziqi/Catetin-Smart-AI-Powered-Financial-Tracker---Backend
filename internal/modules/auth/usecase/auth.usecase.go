package usecase

import (
	"catetin-backend/internal/modules/auth/repository"
)

type AuthUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(r repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{r}
}
