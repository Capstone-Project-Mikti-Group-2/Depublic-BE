package entity

import "time"

type Ticket struct {
	ID        int64      `json:"id"`
	EventID   int64      `json:"event_id"`
	UserID    int64      `json:"user_id"`
	Quantity  int64      `json:"quantity"`
	Total     int64      `json:"total"`
	IsPaid    bool       `json:"is_paid"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempy"`
	BookBy    string     `json:"book_by"`
	UpdateBy  string     `json:"update_by"`
	DeleteBy  string     `json:"delete_by"`
	Event     Event      `json:"event"`
	User      User       `json:"user"`
}

func NewTicket(eventID, userID, quantity int64) *Ticket {
	return &Ticket{
		EventID:   eventID,
		Quantity:  quantity,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}
