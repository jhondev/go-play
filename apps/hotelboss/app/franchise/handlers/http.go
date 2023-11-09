package handlers

import (
	"hotelboss/app"
	cmodels "hotelboss/app/company/models"
	fmodels "hotelboss/app/franchise/models"
	"hotelboss/internal/infra"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request, app *app.App) {
	var input struct {
		Company    cmodels.Company        `json:"company"`
		Franchises []fmodels.FranchiseDTO `json:"franchises"`
	}
	err := app.Read(w, r, &input)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	err = app.Franchise.Create(&input.Company, input.Franchises)
	if err != nil {
		app.ServerError(w, r, err)
		return
	}

	app.Created(w, infra.Envelope{"result": "completed"})
}
