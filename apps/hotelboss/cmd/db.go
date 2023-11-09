package main

import (
	"context"
	"database/sql"
	"hotelboss/internal/infra"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg *config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	// Set the maximum idle timeout.
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// error.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Return the sql.DB connection pool.
	return db, nil
}

func setupDB(db *sql.DB, log *infra.Logger) error {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		log.PrintInfo("db does not need setup", nil)
		return nil
	}

	log.PrintInfo("executing schema.sql", nil)
	sqlBytes, err := os.ReadFile("./data/schema.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return err
	}

	return nil
}
