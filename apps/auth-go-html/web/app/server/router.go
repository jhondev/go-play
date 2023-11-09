package server

import (
	"net/http"
	"umsweb/app"
	"umsweb/app/auth"
	"umsweb/app/user"
	authMiddleware "umsweb/internal/infra/middleware"

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

	// Protected routes
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware.UnloggedInRedirector)

		r.Get("/", app.IndexHandler)
		r.Get("/profile", app.Handler(user.ProfileHandler))
		r.Get("/profile/edit", app.Handler(user.ProfileEditHandler))
		r.Post("/profile/edit", app.Handler(user.ProfileUpdateHandler))
		r.Get("/logout", auth.LogoutHandler)
	})

	// Public routes
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware.LoggedInRedirector)

		r.Get("/login", app.Handler(auth.LoginHandler))
		r.Post("/login", app.Handler(auth.LoginSubmitHandler))

		r.Get("/auth/api/callback", app.Handler(auth.APIAuthCallbackHandler))

		r.Get("/signup", app.Handler(user.SignUpHandler))
		r.Post("/signup", app.Handler(user.SignUpFormHandler))
	})

	return router
}
