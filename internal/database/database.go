package database

import (
	"database/sql"
	"log"
)

const (
	dbHost     = "db"
	dbPort     = 5432
	dbUser     = "Pavel"
	dbPassword = "qwerty"
	dbName     = "demo_service_database"
)

func ConnectDB() *sql.DB {
	// Настройка базы данных PostgreSQL
	db, err := sql.Open("postgres", "postgres://dbUser:dbPassword@db/dbName?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
