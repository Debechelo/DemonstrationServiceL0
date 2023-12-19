package nats

import (
	"DemonstrationServiceL0/internal/service"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func Sender(sc *stan.Conn, done chan bool) {
	for {
		message := service.ReadJSON()

		PublishNatsS(sc, message, done)
		log.Printf("Send message")

		time.Sleep(1 * time.Second)
	}
}
