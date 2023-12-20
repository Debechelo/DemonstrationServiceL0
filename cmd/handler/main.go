package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/database"
	"DemonstrationServiceL0/internal/nats"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

var cfg = config.InitHandler()

func main() {

	//Подключение к базе данных
	db, err := database.ConnectDB(&cfg.DBCfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	done := make(chan bool)
	//Подключение к NATS-Streaming
	sc := nats.ConnectNATSStreaming(&cfg.NATSCfg)
	defer nats.Close(sc)
	fmt.Println("Connected to NATS Streaming")

	//Подключение к Серверу
	//rest.StartServer(":8080")
	//fmt.Println("Connected to Server!")

	// Отправка сообщений
	sub := nats.SubscribeNatsS(&cfg.NATSCfg, db)

	WaitClose(sub, done)

	log.Println("Shutting down...")
}

func WaitClose(sub stan.Subscription, done chan bool) {
	for {
		select {
		case <-done:
			if err := sub.Close(); err != nil {
				log.Println("Closing subscriber:", err)
			}
			break
		}
	}
}
