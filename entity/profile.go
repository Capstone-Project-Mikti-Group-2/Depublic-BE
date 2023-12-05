package entity

import "time"

type Profile struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Address   string    `json:"address"`
	Avatar    []byte    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempy"`
}

func NewProfile(userID int64, address string, avatar []byte) *Profile {
	return &Profile{
		UserID:    userID,
		Address:   address,
		Avatar:    avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateProfile(id int64, address string, avatar []byte) *Profile {
	return &Profile{
		ID:        id,
		Avatar:    avatar,
		Address:   address,
		UpdatedAt: time.Now(),
	}
}
