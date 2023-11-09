package server

import (
	"hotelboss/app"
	"hotelboss/app/franchise/handlers"
	"hotelboss/internal/infra"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(app *app.App) http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	router.Get("/v1/healthcheck", app.Handler(HealthcheckHandler))

	router.Post("/v1/franchises", app.Handler(handlers.Create))

	return router
}

func HealthcheckHandler(w http.ResponseWriter, req *http.Request, app *app.App) {
	data := infra.Envelope{
		"status": "available",
	}

	// Add a temp 4 second delay.
	time.Sleep(2 * time.Second)

	err := app.Success(w, data)
	if err != nil {
		app.ServerError(w, req, err)
	}
}
