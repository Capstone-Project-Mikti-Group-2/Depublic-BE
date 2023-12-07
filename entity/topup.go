package entity

import "time"

type TopUp struct {
	ID        string    `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Amount    int64     `json:"amount" db:"amount"`
	Status    int       `json:"status" db:"status"`
	SnapURL   string    `json:"snap_url" db:"snap_url"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type NewTopup struct {
	ID        string    `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Nominal   int64     `json:"nominal" db:"nominal"`
	Status    int       `json:"status" db:"status"`
	SnapURL   string    `json:"snap_url" db:"snap_url"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
