package app

import (
	"net/http"
	"time"
	authstore "umsapi/app/auth/store"
	userstore "umsapi/app/user/store"
	"umsapi/internal/infra"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/oauth2"
)

type AuthConfig struct {
	ClientRedirectURL string
	GoogleOauth       *oauth2.Config
}
type App struct {
	Jwt  *jwtauth.JWTAuth
	JSON *infra.JSON
	*infra.Errors
	*infra.Logger
	Auth      *AuthConfig
	AuthStore authstore.AuthStore
	UserStore userstore.UserStore
}

func New(jwt *jwtauth.JWTAuth, j *infra.JSON, l *infra.Logger, e *infra.Errors, auth *AuthConfig,
	as authstore.AuthStore, us userstore.UserStore) *App {
	return &App{Jwt: jwt, JSON: j, Errors: e, Logger: l, Auth: auth, AuthStore: as, UserStore: us}
}

func (a *App) Handler(fn func(http.ResponseWriter, *http.Request, *App)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, a)
	}
}

func (app *App) HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	data := infra.Envelope{
		"status": "available",
	}

	// Add a temp 4 second delay.
	time.Sleep(2 * time.Second)

	err := app.JSON.Success(w, data)
	if err != nil {
		app.ServerErrorResponse(w, req, err)
	}
}
