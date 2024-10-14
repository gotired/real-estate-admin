package property

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/gotired/real-estate-admin/server/internal/domain/model/property"
	"github.com/gotired/real-estate-admin/server/internal/domain/repository"
)

type Init struct {
	db *sql.DB
}

func New(db *sql.DB) repository.Property {
	return &Init{
		db: db,
	}
}

func (p *Init) List(skip, limit, order *int, keyword, sort_key, propertyType *string) ([]property.Property, error) {
	orderDirection := "ASC"
	if *order == -1 {
		orderDirection = "DESC"
	}
	conditions := []string{}

	if keyword != nil {
		conditions = append(conditions, fmt.Sprintf("title ~* '%s'", *keyword))
	}
	if propertyType != nil {
		conditions = append(conditions, fmt.Sprintf("type = '%s'", *propertyType))
	}
	filterQuery := ""
	if len(conditions) > 0 {
		filterQuery = fmt.Sprintf("WHERE %s", strings.Join(conditions, " AND "))
	}

	query := fmt.Sprintf(`
		SELECT property_collection.id, property_collection.owner, property_collection.title, property_collection.description, property_collection.type, property_collection.location, property_collection.price, property_collection.photo
		FROM property_collection
		JOIN user_collection ON property_collection.owner = user_collection.id
		%s
		ORDER BY %s %s 
		LIMIT $1 OFFSET $2
		`,
		filterQuery, *sort_key, orderDirection)

	rows, err := p.db.Query(query, limit, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []property.Property
	for rows.Next() {
		var property property.Property
		if err := rows.Scan(&property.ID, &property.Owner, &property.Title, &property.Description, &property.Type, &property.Location, &property.Price, &property.Photo); err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}
	return properties, nil
}

func (p *Init) GetByID(id uuid.UUID) (*property.Property, error) {
	var property property.Property
	err := p.db.QueryRow("SELECT id, owner, title, description, type, location, price , photo FROM property_collection WHERE id = $1", id).Scan(&property.ID, &property.Owner, &property.Title, &property.Description, &property.Type, &property.Location, &property.Price, &property.Photo)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (p *Init) GetByOwnerID(owner uuid.UUID) ([]property.Property, error) {
	rows, err := p.db.Query("SELECT id, owner, title, description, type, location, price, photo FROM property_collection WHERE owner = $1", owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []property.Property
	for rows.Next() {
		var property property.Property
		if err := rows.Scan(&property.ID, &property.Owner, &property.Title, &property.Description, &property.Type, &property.Location, &property.Price, &property.Photo); err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}
	return properties, nil
}

func (p *Init) Create(owner uuid.UUID, title, description, property_type, location, photo string, price int) error {
	_, err := p.db.Exec("INSERT INTO property_collection (id, owner, title, description, type, location, price , photo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", uuid.New(), owner, title, description, property_type, location, price, photo)
	if err != nil {
		return err
	}

	return nil
}

func (p *Init) Delete(id uuid.UUID) error {
	result, err := p.db.Exec("DELETE FROM property_collection where id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows deleted with id %s", id)
	}

	return nil
}

func (p *Init) Update(
	id uuid.UUID,
	owner *uuid.UUID,
	title, description, property_type, location, photo *string,
	price *int,
) error {

	queryParts := []string{}
	args := []interface{}{}
	argIndex := 1

	addQueryPart := func(field string, value interface{}) {
		queryParts = append(queryParts, fmt.Sprintf("%s = $%d", field, argIndex))
		args = append(args, value)
		argIndex++
	}

	if owner != nil {
		addQueryPart("owner", *owner)
	}
	if title != nil {
		addQueryPart("title", *title)
	}
	if description != nil {
		addQueryPart("description", *description)
	}
	if property_type != nil {
		addQueryPart("type", *property_type)
	}
	if location != nil {
		addQueryPart("location", *location)
	}
	if price != nil {
		addQueryPart("price", *price)
	}
	if photo != nil {
		addQueryPart("photo", *photo)
	}

	if len(queryParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE property_collection SET %s WHERE id = $%d",
		strings.Join(queryParts, ", "), argIndex)
	args = append(args, id)

	result, err := p.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated with id %s", id)
	}

	return nil
}
