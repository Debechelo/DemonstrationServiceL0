package main

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/nats"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Подключение к NATS-Streaming
	clientIDSender := "Sender"
	config := config.InitSender()
	var sc *stan.Conn
	go nats.ConnectNATSStreaming(&sc, clientIDSender, config.GetClientID())

	// Ожидайте, пока соединение будет установлено
	for sc == nil {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Connected to NATS Streaming")

	// Отправка сообщений
	go nats.Sender(sc)

	// Ожидание сигналов завершения работы (Ctrl+C)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	log.Println("Shutting down...")

}
