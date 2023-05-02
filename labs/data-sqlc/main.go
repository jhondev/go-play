package main

import (
	"context"
	"database/sql"
	"log"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(ctx)
	log.Print(db)
}
