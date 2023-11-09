package models

import (
	"time"
	cmodels "umsapi/app/common/models"
	"umsapi/internal/validator"

	"github.com/gofrs/uuid"
)

type Profile struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Telephone string    `json:"telephone"`
	External  bool      `json:"external"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *Profile) Validate() *validator.Validator {
	v := validator.New()

	cmodels.ValidateEmail(v, u.Email)

	return v
}
