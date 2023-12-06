package entity

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	OrderID   string    `json:"order_id"`
	UserID    int64     `json:"user_id"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func NewTransaction(orderID string, userID int64, amount int64, status string) *Transaction {
	return &Transaction{
		OrderID: orderID,
		UserID:  userID,
		Amount:  amount,
		Status:  status,
	}
}
