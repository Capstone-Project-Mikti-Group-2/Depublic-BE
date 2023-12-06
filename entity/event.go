package entity

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Image       []byte    `json:"image"`
	Price       int64     `json:"price"`
	Quantity    int64     `json:"quantity"`
	Available   bool      `json:"available"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempy"`
}

func UpdateEvent(id int64, name, description, location string, price, quantity int64, available bool, image []byte, startDate, endDate time.Time) *Event {
	return &Event{
		ID:          id,
		Name:        name,
		Description: description,
		Location:    location,
		Price:       price,
		Quantity:    quantity,
		Available:   available,
		Image:       image,
		StartDate:   startDate,
		EndDate:     endDate,
		UpdatedAt:   time.Now(),
	}
}

func NewEvent(name, description, location string, price, quantity int64, available bool, image []byte, startDate time.Time, endDate time.Time) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		Price:       price,
		Quantity:    quantity,
		Available:   available,
		Image:       image,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
	}
}
