package persistance

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

type UserRespositoryImpl struct {
	db *sql.DB
}

func (r *UserRespositoryImpl) Get(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, name, email, avatar FROM user_collection WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Avatar)
	if err != nil {
		return nil, err
	}

	// Fetch user's properties
	rows, err := r.db.Query("SELECT id, owner, title, description, propertyType, location, price, photo FROM property_collection WHERE owner = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Append properties to user
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

func (r *UserRespositoryImpl) List() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, avatar FROM user_collection")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
