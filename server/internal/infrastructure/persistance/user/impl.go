package user

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
	userSchma "github.com/gotired/real-estate-admin/server/internal/domain/model/user"
	"github.com/gotired/real-estate-admin/server/internal/domain/repository"
)

type UserRespositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) repository.User {
	return &UserRespositoryImpl{
		db: db,
	}
}

func (r *UserRespositoryImpl) Get(id *uuid.UUID) (*userSchma.Base, error) {
	var user userSchma.Base
	err := r.db.QueryRow("SELECT id, name, email, avatar FROM user_collection WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Avatar)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query("SELECT id, owner, title, description, type, location, price, photo FROM property_collection WHERE owner = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []model.Property
	for rows.Next() {
		var property model.Property
		if err := rows.Scan(&property.ID, &property.Owner, &property.Title, &property.Description, &property.Type, &property.Location, &property.Price, &property.Photo); err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}
	user.Property = properties

	return &user, nil
}

func (r *UserRespositoryImpl) List() ([]userSchma.Base, error) {
	rows, err := r.db.Query("SELECT id, name, email, avatar FROM user_collection")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []userSchma.Base
	for rows.Next() {
		var user userSchma.Base
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRespositoryImpl) Create(name *string, email *string, avatar *string) error {
	_, err := r.db.Exec("INSERT INTO user_collection (id, name, email, avatar) VALUES ($1, $2, $3, $4)",
		uuid.New(), name, email, avatar)

	if err != nil {
		return err
	}

	return nil
}
