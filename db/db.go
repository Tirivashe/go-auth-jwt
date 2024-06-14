package db

import (
	"database/sql"
	"log"

	"github.com/Tirivashe/go-fiber-jwt/sqlc"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	connString := "postgres://postgres:shamzio3000@localhost:5432/tutorial?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	DB = db
}

func GetQueries() *sqlc.Queries {
	queries := sqlc.New(DB)
	return queries
}
