package entity

import "time"

type TopUp struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"user_id"`
	Nominal   int64     `json:"nominal"`
	Status    int       `json:"status"`
	SnapURL   string    `json:"snap_url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NewTopup struct {
	UserID    int64     `json:"user_id"`
	Nominal   int64     `json:"nominal"`
	Status    int       `json:"status"`
	SnapURL   string    `json:"snap_url"`
	CreatedAt time.Time `json:"createdAt"`
}
