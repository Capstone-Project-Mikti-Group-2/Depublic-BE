package entity

type AccountBalance struct {
	ID            int64   `json:"id"`
	UserID        int64   `json:"user_id"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}
