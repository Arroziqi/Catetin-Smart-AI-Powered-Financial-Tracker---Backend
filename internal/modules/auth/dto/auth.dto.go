package dto

type RegisterRequest struct {
	Email    string `json:"email" example:"user@mail.com"`
	Password string `json:"password" example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"user@mail.com"`
	Password string `json:"password" example:"123456"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}