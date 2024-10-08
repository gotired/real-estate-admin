package repository

import (
	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/user"
)

type User interface {
	List() ([]user.Base, error)
	Create(name *string, email *string, avatar *string) error
	Get(id *uuid.UUID) (*user.Base, error)
}
