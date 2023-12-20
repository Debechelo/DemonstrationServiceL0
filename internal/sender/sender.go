package sender

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/nats"
	"DemonstrationServiceL0/internal/service"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

func Sender(cfg *config.NATSConfig, done chan bool) {
	json := string(service.ReadJSON())
	for {
		message, id := setID(json)
		nats.PublishNatsS(cfg, []byte(message), done)
		log.Printf("Send message id:%s", id)

		time.Sleep(5 * time.Second)
	}
}

func setID(message string) (string, string) {
	orderUid := uuid.NewString()
	message = strings.Replace(message, "b563feb7b2b84b6test", orderUid, 1)
	return message, orderUid
}
