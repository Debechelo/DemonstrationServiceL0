package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/nats"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {

	done := make(chan bool)
	//Подключение к NATS-Streaming
	clientIDSender := "Sender"
	config := config.InitSender()
	var sc *stan.Conn
	go nats.ConnectNATSStreaming(&sc, clientIDSender, config.GetClientID())

	// Ожидайте, пока соединение будет установлено
	for sc == nil {
		time.Sleep(100 * time.Millisecond)
	}
	defer nats.Close(sc)
	fmt.Println("Connected to NATS Streaming")

	// Отправка сообщений
	go nats.Sender(sc, done)

	<-done

	log.Println("Shutting down...")

}
