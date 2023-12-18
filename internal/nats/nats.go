package nats

import (
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

const (
	natsURL = "nats://nats-streaming:4222"
	subject = "subject"
)

func PublishNatsS(sc *stan.Conn, message []byte) {
	err := (*sc).Publish(subject, message)
	if err != nil {
		log.Fatal(err)
	}
}

func SubscribeNatsS(sc *stan.Conn) stan.Subscription {
	subscription, err := (*sc).Subscribe(subject, func(msg *stan.Msg) {
		// Обработка полученного сообщения
		log.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
	return subscription
}

func Close(sc *stan.Conn) {
	if err := (*sc).Close(); err != nil {
		log.Println("Error closing NATS Streaming connection:", err)
	}
}

func ConnectNATSStreaming(sc **stan.Conn, clientID string, clusterID string) {
	for {
		newSc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL),
			stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
				log.Printf("Connection lost, reason: %v\n", reason)
				ConnectNATSStreaming(sc, clientID, clusterID)
			}),
		)
		if err != nil {
			log.Printf("Error reconnecting: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		*sc = &newSc
		break
	}
}
