package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbHost = "db"
	dbPort = 5432
)

func ConnectDB(dbName string, dbUser string, dbPassword string) *sql.DB {
	// Настройка базы данных PostgreSQL
	dataSourceName := fmt.Sprintf("postgres://%v:%v@db/%v?sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
