package database

import (
	"DemonstrationServiceL0/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func ConnectDB(cfg *config.DBConfig) (*sql.DB, error) {
	// Настройка базы данных PostgreSQL
	var db *sql.DB
	connectionString := getConnectionString(cfg)
	log.Printf(connectionString)
	// Попытки подключения
	reconnected(&db, connectionString)

	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func reconnected(db **sql.DB, connectionString string) {
	var err error
	for {
		*db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Println("Error opening database:", err)
			log.Println("Reconnect in 5 second", err)
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println("Connected to the database!")
			break
		}
	}
}

func getConnectionString(cfg *config.DBConfig) string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.GetDBUser(),
		cfg.GetDBPassword(),
		cfg.GetDBHost(),
		cfg.GetDBPort(),
		cfg.GetDBName())
}
