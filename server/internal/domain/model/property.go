package model

import "github.com/google/uuid"

type Property struct {
	ID          uuid.UUID `json:"id"`
	Owner       uuid.UUID `json:"owner"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Location    string    `json:"location"`
	Price       int       `json:"price"`
	Photo       string    `json:"photo"`
}
