package user

import (
	"errors"
	"net/http"
	"time"
	"umsapi/app"
	cmodels "umsapi/app/common/models"
	"umsapi/app/user/models"
	"umsapi/internal/infra"

	"github.com/go-chi/jwtauth/v5"
	"github.com/gofrs/uuid"
)

// SignUpHandler
func SignUpHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.JSON.Read(w, r, &input)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}
	user := &models.SignUpUser{
		ID:        id,
		Email:     input.Email,
		CreatedAt: time.Now().UTC(),
	}
	err = user.Password.Set(input.Password)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	if v := user.Validate(); !v.Valid() {
		app.FailedValidationResponse(w, r, v.Errors)
		return
	}

	profile, err := app.UserStore.CreateProfile(user)
	if err == cmodels.ErrRecordAlreadyExists {
		app.FailedValidationResponse(w, r, app.Error("conflict", "user is already signed up"))
		return
	}
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	app.JSON.Created(w, infra.Envelope{"profile": profile})
}

// GetProfileHandler
func GetProfileHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	id, err := getUserIDFromToken(r)
	if err != nil {
		app.FailedValidationResponse(w, r, app.Error("user_id", err.Error()))
	}

	profile, err := app.UserStore.GetProfile(id)
	if err == cmodels.ErrRecordNotFound {
		app.FailedValidationResponse(w, r, app.Error("not_found", "profile not found"))
		return
	}
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	app.JSON.Success(w, infra.Envelope{"profile": profile})
}

// UpdateProfileHandler
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	id, err := getUserIDFromToken(r)
	if err != nil {
		app.FailedValidationResponse(w, r, app.Error("user_id", err.Error()))
	}

	// GET ENTIRE PROFILE TO BE PATCHED
	profile, err := app.UserStore.GetProfile(id)
	if err == cmodels.ErrRecordNotFound {
		app.FailedValidationResponse(w, r, app.Error("not_found", "profile not found"))
		return
	}
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	var input struct {
		Email     *string `json:"email"`
		Name      *string `json:"name"`
		Telephone *string `json:"telephone"`
		External  *string `json:"external"`
	}
	err = app.JSON.Read(w, r, &input)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	// PATCH PROFILE
	if input.Email != nil && profile.Email != *input.Email {
		profile.Email = *input.Email
		p, _ := app.UserStore.GetProfileByEmail(profile.Email)
		if p != nil {
			app.FailedValidationResponse(w, r, app.Error("conflict", "new email is already registered"))
			return
		}
	}
	if input.Name != nil && profile.Name != *input.Name {
		profile.Name = *input.Name
	}
	if input.Telephone != nil && profile.Telephone != *input.Telephone {
		profile.Telephone = *input.Telephone
	}

	if v := profile.Validate(); !v.Valid() {
		app.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.UserStore.UpdateProfile(profile)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}
}

func getUserIDFromToken(r *http.Request) (uuid.UUID, error) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userID, ok := claims["user_id"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("no user_id claim found in token")
	}
	id, err := uuid.FromString(userID)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid user_id in token")
	}
	return id, nil
}
