package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

func DSNFromEnv() string {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "db"
	}
	return fmt.Sprintf("postgres://%s:%s@%s/%s", user, pass, host, db)
}

func Connect(dsn string) (*sql.DB, error) {
	return sql.Open("pgx", dsn)
}
