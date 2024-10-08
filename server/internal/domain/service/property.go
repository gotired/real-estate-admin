package service

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
	"github.com/gotired/real-estate-admin/server/internal/domain/repository"
)

type Property struct {
	Repository repository.Property
}

func (s *Property) List() ([]model.Property, error) {
	return s.Repository.List()
}

func (s *Property) Get(id *uuid.UUID) (*model.Property, error) {
	return s.Repository.Get(id)
}

func (s *Property) Create(id *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error {
	return s.Repository.Create(id, title, description, property_type, location, price, photo)
}

func (s *Property) Delete(id *uuid.UUID) error {
	return s.Repository.Delete(id)
}

func (s *Property) Update(id *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error {
	return s.Repository.Update(id, title, description, property_type, location, price, photo)
}
