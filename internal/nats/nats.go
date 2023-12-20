package nats

import (
	"DemonstrationServiceL0/internal/config"
	"DemonstrationServiceL0/internal/handler"
	"database/sql"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

var sc *stan.Conn

func PublishNatsS(cfg *config.NATSConfig, message []byte, done chan bool) {
	err := (*sc).Publish(cfg.GetSubject(), message)
	if err != nil {
		log.Printf("Received message: %v\n", err)
		if !(*sc).NatsConn().IsClosed() {
			log.Printf("close channel")
			done <- true
		}
		log.Printf("Wait 2 second")
		time.Sleep(2 * time.Second)
		PublishNatsS(cfg, message, done)
	}
}

func SubscribeNatsS(cfg *config.NATSConfig, db *sql.DB) stan.Subscription {
	subscription, err := (*sc).Subscribe(cfg.GetSubject(), func(msg *stan.Msg) {
		handler.Handler(msg, db)
	}, stan.DurableName("Handler"))
	if err != nil {
		log.Printf("Subscription error%v:", err)
	}
	log.Printf("Subscription")
	return subscription
}

func Close(sc *stan.Conn) {
	if err := (*sc).Close(); err != nil {
		log.Println("Error closing NATS Streaming connection:", err)
	}
}

func ConnectNATSStreaming(cfg *config.NATSConfig) *stan.Conn {
	for {
		newSc, err := stan.Connect(cfg.GetClusterID(), cfg.GetClientID(), stan.NatsURL(cfg.GetUrl()),
			stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
				log.Printf("Connection lost, reason: %v\n", reason)
				sc = ConnectNATSStreaming(cfg)
			}),
		)
		if err != nil {
			log.Printf("Error reconnecting: %v. Retrying in 5 seconds...", err, err)
			time.Sleep(5 * time.Second)
			continue
		}
		sc = &newSc
		break
	}

	return sc
}
