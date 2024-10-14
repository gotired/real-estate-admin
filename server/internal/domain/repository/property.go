package repository

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/property"
)

type Property interface {
	List(skip, limit, order *int, keyword, sort_key, propertyType *string) ([]property.Property, error)
	GetByID(id uuid.UUID) (*property.Property, error)
	GetByOwnerID(owner uuid.UUID) ([]property.Property, error)
	Create(owner uuid.UUID, title, description, propertyType, location, photo string, price int) error
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, owner *uuid.UUID, title, description, propertyType, location, photo *string, price *int) error
}
