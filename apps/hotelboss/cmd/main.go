package main

import (
	"hotelboss/app"
	cstore "hotelboss/app/company/store"
	"hotelboss/app/franchise/core"
	"hotelboss/app/franchise/store"
	"hotelboss/app/scraper"
	"hotelboss/internal/infra"
	"hotelboss/internal/server"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	cfg := getConfig()

	logger := infra.NewLogger(os.Stdout, infra.LevelInfo)
	errors := infra.NewErrors(logger)
	resp := infra.NewResponse(logger, errors)
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	err = setupDB(db, logger)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	scraper := scraper.New()
	store := store.New(db)
	cstore := cstore.New(db)
	franchise := core.New(scraper, store, cstore)
	app := app.New(franchise, logger, errors, resp)

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
