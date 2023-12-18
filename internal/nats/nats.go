package nats

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

const (
	clusterID = "your-cluster-id"
	clientID  = "your-client-id"
	natsURL   = "nats://nats-streaming:4222"
)

func ConnectNatsS() stan.Conn {
	//Настройка nats-streaming
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
	}
	return sc
}

func PublishNatsS(sc stan.Conn) {
	err := sc.Publish("your-channel", []byte("your-message"))
	if err != nil {
		log.Fatal(err)
	}
}

func SubscribeNatsS(sc stan.Conn) stan.Subscription {
	subscription, err := sc.Subscribe("your-channel", func(msg *stan.Msg) {
		// Обработка полученного сообщения
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
	return subscription
}

func Close(sc stan.Conn) {
	if err := sc.Close(); err != nil {
		log.Println("Error closing NATS Streaming connection:", err)
	}
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("Connection lost, reason: %v\n", reason)
			//

		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
