package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/database"
	"DemonstrationServiceL0/internal/nats"
	"DemonstrationServiceL0/internal/transport/rest"
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

	//Подключение к Серверу
	rest.StartServer(":8080")
	fmt.Println("Connected to Server!")

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

	// Отправка сообщений

	<-done

	log.Println("Shutting down...")

}
