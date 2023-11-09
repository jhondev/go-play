package server

import (
	"net/http"
	"umsapi/app"
	"umsapi/app/auth"
	"umsapi/app/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func NewRouter(app *app.App) http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		jwtauth.Verifier(app.Jwt), // Seek, verify and validate JWT tokens
	)

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Authenticator)

		r.Get("/profile", app.Handler(user.GetProfileHandler))
		r.Patch("/profile", app.Handler(user.UpdateProfileHandler))
	})

	// public routes
	router.Group(func(r chi.Router) {
		r.Get("/healthcheck", app.HealthcheckHandler)
		r.Post("/login", app.Handler(auth.LoginHandler))
		r.Get("/auth/google/login", app.Handler(auth.GoogleAuthHandler))
		r.Get("/auth/google/callback", app.Handler(auth.GoogleAuthCallbackHandler))
		r.Post("/signup", app.Handler(user.SignUpHandler))
	})

	return router
}
