package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Avatar   string     `json:"avatar"`
	Property []Property `json:"property"`
}
