package user

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

type Base struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	Email    string           `json:"email"`
	Avatar   string           `json:"avatar"`
	Property []model.Property `json:"property"`
}
