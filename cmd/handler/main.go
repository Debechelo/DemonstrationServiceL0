package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/database"
	"DemonstrationServiceL0/internal/nats"
	"DemonstrationServiceL0/internal/transport/rest"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	clientIDHandler := "Handler"
	var sc *stan.Conn
	go nats.ConnectNATSStreaming(&sc, clientIDHandler, config.GetClientID())

	// Ожидайте, пока соединение будет установлено
	for sc == nil {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Connected to NATS Streaming")

	// Обработка сообщений

	// Ожидание сигналов завершения работы (Ctrl+C)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	log.Println("Shutting down...")

}
