package app

import (
	"hotelboss/app/franchise/core"
	"hotelboss/internal/infra"
	"net/http"
)

type App struct {
	Franchise *core.Franchise
	*infra.Logger
	*infra.Errors
	*infra.Response
}

func New(
	f *core.Franchise,
	l *infra.Logger,
	e *infra.Errors,
	resp *infra.Response) *App {
	return &App{Franchise: f, Logger: l, Errors: e, Response: resp}
}

func (a *App) Handler(fn func(http.ResponseWriter, *http.Request, *App)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, a)
	}
}
