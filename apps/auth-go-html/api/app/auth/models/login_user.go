package models

import (
	cmodels "umsapi/app/common/models"
	"umsapi/internal/validator"
)

type LoginUser struct {
	Email    string
	Password cmodels.Password
}

func (u *LoginUser) Validate() *validator.Validator {
	v := validator.New()

	cmodels.ValidateEmail(v, u.Email)
	cmodels.ValidatePasswordPlaintext(v, *u.Password.Plaintext)

	return v
}
