package repository

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

type Property interface {
	List() ([]model.Property, error)
	Get(id *uuid.UUID) (*model.Property, error)
	Create(id *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error
	Delete(id *uuid.UUID) error
	Update(id *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error
}
