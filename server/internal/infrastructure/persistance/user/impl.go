package user

import (
	"database/sql"

	"github.com/google/uuid"
	userSchema "github.com/gotired/real-estate-admin/server/internal/domain/model/user"
	"github.com/gotired/real-estate-admin/server/internal/domain/repository"
)

type initialize struct {
	db *sql.DB
}

func New(db *sql.DB) repository.User {
	return &initialize{
		db: db,
	}
}

func (r *initialize) Get(id uuid.UUID) (*userSchema.Base, error) {
	var user userSchema.Base
	err := r.db.QueryRow("SELECT id, name, email, avatar FROM user_collection WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Avatar)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *initialize) List(skip *int, limit *int) ([]userSchema.Base, error) {
	rows, err := r.db.Query("SELECT id, name, email, avatar FROM user_collection LIMIT $1 OFFSET $2", limit, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []userSchema.Base
	for rows.Next() {
		var user userSchema.Base
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Avatar); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *initialize) Create(name string, email string, avatar string) error {
	_, err := r.db.Exec("INSERT INTO user_collection (id, name, email, avatar) VALUES ($1, $2, $3, $4)",
		uuid.New(), name, email, avatar)

	if err != nil {
		return err
	}

	return nil
}

func (r *initialize) GetByEmail(email string) (*uuid.UUID, error) {
	var user userSchema.Base
	err := r.db.QueryRow("SELECT id, name, email, avatar FROM user_collection WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Avatar)
	if err != nil {
		return nil, err
	}
	return &user.ID, nil
}
