package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/nats"
	"DemonstrationServiceL0/internal/sender"
	"fmt"
	"log"
)

var cfg = config.InitSender()

func main() {

	done := make(chan bool)
	//Подключение к NATS-Streaming
	sc := nats.ConnectNATSStreaming(&cfg.NATSCfg)
	defer nats.Close(sc)
	fmt.Println("Connected to NATS Streaming")

	// Отправка сообщений
	go sender.Sender(&cfg.NATSCfg, done)

	WaitClose(done)

	log.Println("Shutting down...")
}

func WaitClose(done chan bool) {
	for {
		select {
		case <-done:
			break
		}
	}
}
