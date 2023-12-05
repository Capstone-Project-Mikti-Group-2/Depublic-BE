package entity

type AccountBalance struct {
	ID            int64   `db:"id"`
	UserID        int64   `db:"user_id"`
	AccountNumber string  `db:"account_number"`
	Balance       float64 `db:"balance"`
}
