package auth

import (
	"net/http"
	"umsapi/app"
	"umsapi/app/auth/models"
	"umsapi/internal/infra"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	// Create an anonymous struct to hold the expected data from the request body.
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.JSON.Read(w, r, &input)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	user := &models.LoginUser{Email: input.Email}
	err = user.Password.Set(input.Password)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	if v := user.Validate(); !v.Valid() {
		app.FailedValidationResponse(w, r, v.Errors)
		return
	}

	valid, id, err := app.AuthStore.Login(user.Email, user.Password)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}
	if !valid {
		app.FailedValidationResponse(w, r, app.Error("login", "invalid email or password"))
		return
	}

	app.JSON.Success(w, infra.Envelope{"id": id})
}
