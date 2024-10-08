package property

import (
	"database/sql"
)

type PropertysRespositoryImpl struct {
	db *sql.DB
}

// func New(db *sql.DB) repository.Property {
// 	return &PropertysRespositoryImpl{
// 		db: db,
// 	}
// }

// func (p *PropertysRespositoryImpl) List() ([]model.Property, error) {
// 	return nil
// }

// func (p *PropertysRespositoryImpl) Get(id *uuid.UUID) (*model.Property, error) {
// 	return nil
// }

// func (p *PropertysRespositoryImpl) Create(id *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error {
// 	return nil
// }

// func (p *PropertysRespositoryImpl) Delete(id *uuid.UUID) error {
// 	return nil
// }

// func (p *PropertysRespositoryImpl) Update(id *uuid.UUID, title *string, description *string, property_type *string, location *string, price *int, photo *string) error {
// 	return nil
// }
