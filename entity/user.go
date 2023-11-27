package entity

import (
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"username"`
	Email     string     `json:"email"`
	Number    string     `json:"number"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempy"`
}

func NewUser(name, email, number, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Number:    number,
		CreatedAt: time.Now(),
	}
}

func UpdateUser(id int64, name, email, number, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}

func Register(name, email, password, number string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Number:   number,
	}
}
