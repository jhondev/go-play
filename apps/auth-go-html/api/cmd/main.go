package main

import (
	"log"
	"net/http"
	"os"
	"umsapi/app"
	authstore "umsapi/app/auth/store"
	"umsapi/app/server"
	userstore "umsapi/app/user/store"
	"umsapi/internal/infra"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	cfg := getConfig()

	jwt := jwtauth.New("HS256", []byte(cfg.jwtsecret), nil)
	json := infra.NewJSON()
	logger := infra.NewLogger(os.Stdout, infra.LevelInfo)
	errors := infra.NewErrors(logger, json)
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	auth := &app.AuthConfig{
		ClientRedirectURL: cfg.auth.clientRedirectURL,
		GoogleOauth:       cfg.googleOauthConfig,
	}
	userStore := userstore.New(db)
	authStore := authstore.New(db)
	app := app.New(jwt, json, logger, errors, auth, authStore, userStore)

	router := server.NewRouter(app)

	address := ":" + cfg.port
	log.Println("serving on " + address)
	err = http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(router, &http2.Server{}),
	)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
