package app

import (
	"net/http"
	"umsweb/internal/client"
	"umsweb/internal/infra/template"

	"github.com/go-chi/jwtauth/v5"
)

type App struct {
	Tmplr           template.Templater
	Jwt             *jwtauth.JWTAuth
	Client          *client.APIClient
	AuthExternalURL string
}

func New(tmplr template.Templater, jwt *jwtauth.JWTAuth, client *client.APIClient) *App {
	return &App{Tmplr: tmplr, Jwt: jwt, Client: client}
}

func (a *App) Handler(fn func(http.ResponseWriter, *http.Request, *App)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, a)
	}
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/profile", http.StatusFound)
}
