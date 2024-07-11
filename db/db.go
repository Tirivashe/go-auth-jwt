package db

import (
	"database/sql"
	"log"

	"github.com/Tirivashe/go-fiber-jwt/sqlc"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	connString := "postgresql://tutorial_k0uh_user:rP8wrRcuITtkhlMfCAJQx7sgKqlTBUxk@dpg-cq7qq3bv2p9s73c9b1h0-a.ohio-postgres.render.com/tutorial_k0uh"
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
