package service

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/user"
	"github.com/gotired/real-estate-admin/server/internal/domain/repository"
)

type User struct {
	Repository repository.User
}

func (s *User) List() ([]user.Base, error) {
	return s.Repository.List()
}

func (s *User) Create(name *string, email *string, avatar *string) error {
	return s.Repository.Create(name, email, avatar)
}

func (s *User) Get(id *uuid.UUID) (*user.Base, error) {
	return s.Repository.Get(id)
}
