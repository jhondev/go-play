package store

import (
	"context"
	"database/sql"
	"time"
	cmodels "umsapi/app/common/models"
	"umsapi/app/user/models"

	"github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
)

const alreadyExists = 1062

type UserStore interface {
	GetProfile(id uuid.UUID) (*models.Profile, error)
	GetProfileByEmail(email string) (*models.Profile, error)
	CreateProfile(user *models.SignUpUser) (*models.Profile, error)
	UpdateProfile(profile *models.Profile) error
}

func New(db *sql.DB) UserStore {
	return &userStore{db: db}
}

type userStore struct {
	db *sql.DB
}

// GetProfile
func (s *userStore) GetProfile(id uuid.UUID) (*models.Profile, error) {
	query := "SELECT id, email, password_hash, name, telephone, created_at, updated_at FROM users WHERE id = ?"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := s.db.QueryRowContext(ctx, query, id)
	return scanProfile(row)
}

// GetProfileByEmail
func (s *userStore) GetProfileByEmail(email string) (*models.Profile, error) {
	query := "SELECT id, email, password_hash, name, telephone, created_at, updated_at FROM users WHERE email = ?"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := s.db.QueryRowContext(ctx, query, email)
	return scanProfile(row)
}

func scanProfile(row *sql.Row) (*models.Profile, error) {
	var profile models.Profile
	var password_hash []byte
	err := row.Scan(&profile.ID, &profile.Email, &password_hash, &profile.Name, &profile.Telephone, &profile.CreatedAt, &profile.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, cmodels.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	profile.External = password_hash == nil

	return &profile, nil
}

// CreateProfile
func (s *userStore) CreateProfile(user *models.SignUpUser) (*models.Profile, error) {
	cmd := "INSERT INTO users (id, email, name, password_hash, telephone, created_at) VALUES (?,?,?,?,?,?)"

	profile := &models.Profile{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Telephone: "",
		CreatedAt: user.CreatedAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []any{user.ID, user.Email, user.Name, user.Password.Hash, profile.Telephone, user.CreatedAt}
	_, err := s.db.ExecContext(ctx, cmd, args...)

	if err != nil {
		if merr := err.(*mysql.MySQLError); merr != nil && merr.Number == alreadyExists {
			return nil, cmodels.ErrRecordAlreadyExists
		}

		return nil, err
	}

	return profile, nil
}

// UpdateProfile
func (s *userStore) UpdateProfile(profile *models.Profile) error {
	cmd := "UPDATE users SET email = ?, name = ?, telephone = ? WHERE id = ?"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []any{profile.Email, profile.Name, profile.Telephone, profile.ID}
	_, err := s.db.ExecContext(ctx, cmd, args...)
	if err != nil {
		return err
	}

	return nil
}
