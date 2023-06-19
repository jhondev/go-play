package main

import (
	"context"
	"log"
	"sqlc-test/data"
	"sqlc-test/games"

	"github.com/gertd/go-pluralize"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pocketbase/dbx"
)

func main() {
	db, err := dbx.Open("pgx", "postgres://postgres:postgres@/arena")
	if err != nil {
		log.Fatal(err)
	}
	p := pluralize.NewClient()
	db.TableMapper = func(a interface{}) string {
		return p.Plural(dbx.GetTableName(a))
	}
	games := games.New(db)
	// server := grpc.New(games)
	err = games.CreateAll(context.Background(), &data.Game{
		KeyName: "lol",
		Name:    "league of legends",
	})
	if err != nil {
		panic(err)
	}

	log.Println("all good")
}
