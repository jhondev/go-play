// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package data

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type Game struct {
	ID      int64
	KeyName string
	Name    string
}

type Match struct {
	ID    int64
	Name  string
	Score sql.NullString
}
