package store

import (
	"database/sql"
	"hotelboss/app/franchise/models"
)

type FranchiseStore interface {
	Create(f *models.NewFranchise)
	Update(f *models.NewFranchise)
	Query()
}

func New(db *sql.DB) FranchiseStore {
	return &store{db: db}
}

type store struct {
	db *sql.DB
}

// Create implements FranchiseStore.
func (*store) Create(f *models.NewFranchise) {
	panic("unimplemented")
}

// Query implements FranchiseStore.
func (*store) Query() {
	panic("unimplemented")
}

// Update implements FranchiseStore.
func (*store) Update(f *models.NewFranchise) {
	panic("unimplemented")
}
