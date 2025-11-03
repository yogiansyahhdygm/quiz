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
	dsn := os.Getenv("DATABASE_URL") // contoh: postgres://user:pass@host:port/dbname
	// dsn := "postgres://postgres:admin@localhost:5432/perpustakaan?sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database tidak bisa diakses:", err)
	}

	fmt.Println("Database connected!")
}
