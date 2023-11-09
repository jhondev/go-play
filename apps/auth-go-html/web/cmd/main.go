package main

import (
	"log"
	"net/http"
	"umsweb/app"
	"umsweb/app/server"
	"umsweb/internal/client"
	"umsweb/internal/infra/template"

	"github.com/go-chi/jwtauth/v5"
)

func main() {
	cfg := getConfig()

	router := server.NewRouter(&app.App{
		Tmplr:           template.NewTemplater("tmpl/*.html"),
		Jwt:             jwtauth.New("HS256", []byte(cfg.jwtsecret), nil),
		Client:          client.NewAPIClient(cfg.apiURL),
		AuthExternalURL: cfg.authExternalURL,
	})

	address := ":" + cfg.port
	log.Println("serving on " + address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}
