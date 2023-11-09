package models

import comodels "hotelboss/app/common/models"

type Contact struct {
	Email    string            `json:"email"`
	Phone    string            `json:"phone"`
	Location comodels.Location `json:"location"`
}

type Owner struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Contact   Contact `json:"contact"`
}

type Informacion struct {
	Name      string            `json:"name"`
	TaxNumber string            `json:"tax_number"`
	Location  comodels.Location `json:"location"`
}

type Company struct {
	Owner       Owner       `json:"owner"`
	Informacion Informacion `json:"informacion"`
}
