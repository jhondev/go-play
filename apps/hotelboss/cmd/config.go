package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	port string
	db   *dbConfig
}

func getConfig() *config {
	godotenv.Load()

	return &config{
		port: getEnv("PORT", "8081"),
		db: &dbConfig{
			dsn:          getEnv("DB_DSN", "./data/hotelboss.db"),
			maxOpenConns: atoi(getEnv("DB_MAX_OPEN_CONNS", "25")),
			maxIdleConns: atoi(getEnv("DB_MAX_IDLE_CONNS", "25")),
			maxIdleTime:  getEnv("DB_MAX_IDLE_TIME", "15m"),
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

// initial config needs to be correct so panic is ok
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
