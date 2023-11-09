package main

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	port            string
	apiURL          string
	jwtsecret       string
	authExternalURL string
}

func getConfig() *config {
	godotenv.Load()

	cfg := &config{
		port:      getEnv("PORT", "8080"),
		apiURL:    getEnv("API_URL", ""),
		jwtsecret: getEnv("JWT_SECRET", ""),
	}
	cfg.authExternalURL = cfg.apiURL + "/auth/google/login"

	return cfg
}

// getEnv gets a value from an environment variable (returns 'value' if the variable isn't found)
func getEnv(key string, value string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return value
}
