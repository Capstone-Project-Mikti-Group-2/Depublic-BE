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
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

func UpdateEvent(id int64, name, description, location string, price, quantity int64, available bool, image []byte, startDate string, endDate string) *Event {
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

func NewEvent(id int64, name, description, location string, price, quantity int64, available bool, image []byte, startDate string, endDate string) *Event {
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
		CreatedAt:   time.Now(),
	}
}
