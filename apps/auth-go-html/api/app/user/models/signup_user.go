package models

import (
	"time"
	cmodels "umsapi/app/common/models"
	"umsapi/internal/validator"

	"github.com/gofrs/uuid"
)

type SignUpUser struct {
	ID        uuid.UUID
	Email     string
	Name      string
	Password  cmodels.Password
	CreatedAt time.Time
}

func (u *SignUpUser) Validate() *validator.Validator {
	v := validator.New()

	cmodels.ValidateEmail(v, u.Email)
	cmodels.ValidatePasswordPlaintext(v, *u.Password.Plaintext)

	return v
}
