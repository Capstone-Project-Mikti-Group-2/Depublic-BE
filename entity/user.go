package entity

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	Number    string    `json:"number"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	Profile   *Profile  `json:"profile"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func NewUser(name, email, number, password, role string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Number:    number,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateUser(id int64, name, email, number, password, role string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Number:    number,
		Role:      role,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}

func Register(name, email, password, number, role string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Number:   number,
		Role:     role,
	}
}

//final entity user
