package models

import (
	comodels "hotelboss/app/common/models"
)

type FranchiseDTO struct {
	Name     string            `json:"name"`
	URL      string            `json:"url"`
	Location comodels.Location `json:"location"`
}
