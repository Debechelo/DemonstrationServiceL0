package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/database"
	"DemonstrationServiceL0/internal/nats"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {

	config := config.InitHandler()

	//Подключение к базе данных
	db := database.ConnectDB(config.GetDBName(), config.GetDBUser(), config.GetDBPassword())
	defer db.Close()

	//Подключение к NATS-Streaming
	done := make(chan bool)
	clientIDHandler := "Handler"
	var sc *stan.Conn
	go nats.ConnectNATSStreaming(&sc, clientIDHandler, config.GetClientID())

	// Ожидайте, пока соединение будет установлено
	for sc == nil {
		time.Sleep(100 * time.Millisecond)
	}
	defer nats.Close(sc)
	fmt.Println("Connected to NATS Streaming")

	//Подключение к Серверу
	//rest.StartServer(":8080")
	//fmt.Println("Connected to Server!")

	// Отправка сообщений
	sub := nats.SubscribeNatsS(sc, clientIDHandler, db)

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
