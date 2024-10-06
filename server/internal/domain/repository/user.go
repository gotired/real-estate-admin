package repository

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

type User interface {
	List() ([]model.User, error)
	Create(name string, email string, avatar string) error
	Get(id uuid.UUID) (model.User, error)
}
