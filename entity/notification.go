package entity

import "time"

type Notification struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewNotification(id int, typ string, content string, isRead bool, createdAt time.Time) Notification {
	return Notification{
		ID:        id,
		Type:      typ,
		Content:   content,
		IsRead:    isRead,
		CreatedAt: createdAt,
	}
}

func HistoryNotification(id int, typ string, content string, isRead bool, createdAt time.Time) Notification {
	return Notification{
		ID:      id,
		Type:    typ,
		Content: content,
		IsRead:  isRead,
	}
}
