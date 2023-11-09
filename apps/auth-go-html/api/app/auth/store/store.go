package store

import (
	"context"
	"database/sql"
	"time"
	"umsapi/app/common/models"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthStore interface {
	Login(email string, password models.Password) (bool, uuid.UUID, error)
}

func New(db *sql.DB) AuthStore {
	return &authStore{db: db}
}

type authStore struct {
	db *sql.DB
}

// Login
func (s *authStore) Login(email string, password models.Password) (bool, uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id uuid.UUID
	var phash []byte
	err := s.db.QueryRowContext(ctx, "SELECT id, password_hash FROM users WHERE email = ?", email).Scan(&id, &phash)
	if err == sql.ErrNoRows {
		return false, uuid.UUID{}, nil
	}
	if err != nil {
		return false, uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword(phash, []byte(*password.Plaintext))
	if err != nil {
		return false, uuid.UUID{}, nil
	}

	return true, id, nil
}
