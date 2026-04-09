package usecase

import (
	"catetin-backend/internal/modules/auth/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *AuthUsecase) Register(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := &domain.User{
		Email:    email,
		Password: string(hashed),
	}

	return u.repo.Create(user)
}
