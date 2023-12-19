package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	dbHost = "db"
	dbPort = 5432
)

func ConnectDB(dbName string, dbUser string, dbPassword string) *sql.DB {
	// Настройка базы данных PostgreSQL
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var db *sql.DB
	var err error

	// Попытки подключения
	for {
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Println("Error opening database:", err)
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Connected to the database!")
			break
		}
	}
	return db
}
