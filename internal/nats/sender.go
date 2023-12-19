package nats

import (
	"DemonstrationServiceL0/internal/service"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"strings"
	"time"
)

func Sender(sc *stan.Conn, done chan bool) {
	json := string(service.ReadJSON())
	for {
		message, id := setID(json)
		PublishNatsS(sc, []byte(message), done)
		log.Printf("Send message id:%s", id)

		time.Sleep(1 * time.Second)
	}
}

func setID(message string) (string, string) {
	orderUid := uuid.NewString()
	message = strings.Replace(message, "b563feb7b2b84b6test", orderUid, 1)
	return message, orderUid
}
