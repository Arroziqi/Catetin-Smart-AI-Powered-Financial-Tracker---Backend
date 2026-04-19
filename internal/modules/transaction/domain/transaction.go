package domain

import (
	"time"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"` // income | expense
	Amount    int       `json:"amount"`
	Category  string    `json:"category"`
	Note      string    `json:"note"`
	Date      string    `json:"date"` // YYYY-MM-DD
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
