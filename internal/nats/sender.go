package nats

import (
	"DemonstrationServiceL0/internal/service"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func Sender(sc *stan.Conn) {
	for {
		message := service.ReadJSON()

		PublishNatsS(sc, message)
		log.Printf("Send message")

		time.Sleep(1 * time.Second)
	}
}
