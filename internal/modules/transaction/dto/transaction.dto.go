package dto

type CreateTransactionRequest struct {
	Type     string `json:"type" validate:"required"` // income | expense
	Amount   int    `json:"amount" validate:"required,min=1"`
	Category string `json:"category" validate:"required"`
	Note     string `json:"note"`
	Date     string `json:"date" validate:"required"` // YYYY-MM-DD
}

type UpdateTransactionRequest struct {
	Type     string `json:"type" validate:"required"`
	Amount   int    `json:"amount" validate:"required,min=1"`
	Category string `json:"category" validate:"required"`
	Note     string `json:"note"`
	Date     string `json:"date" validate:"required"` // YYYY-MM-DD
}

type GetTransactionQuery struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Category  string `query:"category"`
	Type      string `query:"type"`
	StartDate string `query:"start_date"`
	EndDate   string `query:"end_date"`
}
