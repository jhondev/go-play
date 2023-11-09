package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type dbConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}
type authConfig struct {
	clientRedirectURL string
}
type config struct {
	port              string
	jwtsecret         string
	db                *dbConfig
	auth              *authConfig
	googleOauthConfig *oauth2.Config
}

func getConfig() *config {
	godotenv.Load()

	// initial config needs to be correct so panic is ok
	atoi := func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}

	return &config{
		port:      getEnv("PORT", "8081"),
		jwtsecret: getEnv("JWT_SECRET", ""),
		db: &dbConfig{
			dsn:          getEnv("DB_DSN", ""),
			maxOpenConns: atoi(getEnv("DB_MAX_OPEN_CONNS", "25")),
			maxIdleConns: atoi(getEnv("DB_MAX_IDLE_CONNS", "25")),
			maxIdleTime:  getEnv("DB_MAX_IDLE_TIME", "15m"),
		},
		auth: &authConfig{
			clientRedirectURL: getEnv("AUTH_CLIENT_REDIRECT_URL", ""),
		},
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  getEnv("AUTH_GOOGLE_REDIRECT_URL", ""),
			ClientID:     getEnv("AUTH_GOOGLE_CLIENT_ID", ""),
			ClientSecret: getEnv("AUTH_GOOGLE_CLIENT_SECRET", ""),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

// getEnv gets a value from an environment variable (returns 'value' if the variable isn't found)
func getEnv(key string, value string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return value
}
