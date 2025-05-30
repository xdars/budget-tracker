package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func SetupPostgresConn() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env\n")
		return nil
	}

	connString := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", connString)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Postgres connection OK")

	return db
}
