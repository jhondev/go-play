package store

import (
	"database/sql"
	"hotelboss/app/company/models"
)

type CompanyStore interface {
	Create(c *models.Company) error
}

func New(db *sql.DB) CompanyStore {
	return &store{db: db}
}

type store struct {
	db *sql.DB
}

// Create implements CompanyStore.
func (s *store) Create(c *models.Company) error {
	_, err := s.db.Exec(
		`INSERT INTO Owners (Email,FirstName,LastName,Phone,LocationID) VALUES (?,?,?,?,?);
		INSERT INTO Locations (ID,City,Country,Address,ZipCode) VALUES (?,?,?,?,?);
		INSERT INTO companies (TaxNumber,OwnerID,Name,LocationID) VALUES (?,?,?,?)`,
		c.Owner.Contact.Email,
		c.Owner.FirstName,
		c.Owner.LastName,
		c.Owner.Contact.Phone,
		c.Owner.Contact.Location.ZipCode, // TODO: validate ids for locations
		c.Owner.Contact.Location.ZipCode, // TODO: validate ids for locations
		c.Owner.Contact.Location.City,
		c.Owner.Contact.Location.Country,
		c.Owner.Contact.Location.Address,
		c.Owner.Contact.Location.ZipCode,
		c.Informacion.TaxNumber,
		c.Owner.Contact.Email,
		c.Informacion.Name,
		c.Owner.Contact.Location.ZipCode) // TODO: validate ids for locations

	return err
}
