package service

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/property"
	"github.com/gotired/real-estate-admin/server/internal/domain/repository"
)

type Property struct {
	Repository repository.Property
}

func (s *Property) List(skip, limit, order *int, keyword, sort_keyword, propertyType *string) ([]property.Property, error) {
	return s.Repository.List(skip, limit, order, keyword, sort_keyword, propertyType)
}

func (s *Property) GetByID(id uuid.UUID) (*property.Property, error) {
	return s.Repository.GetByID(id)
}

func (s *Property) GetByOwnerID(owner uuid.UUID) ([]property.Property, error) {
	return s.Repository.GetByOwnerID(owner)
}

func (s *Property) Create(owner uuid.UUID, title, description, propertyType, location, photo string, price int) error {
	return s.Repository.Create(owner, title, description, propertyType, location, photo, price)
}

func (s *Property) Delete(id uuid.UUID) error {
	return s.Repository.Delete(id)
}

func (s *Property) Update(id uuid.UUID, owner *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error {
	return s.Repository.Update(id, owner, title, description, property_type, location, photo, price)
}
