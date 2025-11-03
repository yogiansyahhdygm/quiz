package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not found")
	}

	fmt.Println("Database URL:", dbURL)

	dbURL += "?sslmode=require"
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database tidak bisa diakses:", err)
	}

	fmt.Println("Database connected!")

}
