package configs

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDBCon() *pgxpool.Pool {
	DSN := os.Getenv("DSN")

	db, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	return db
}
