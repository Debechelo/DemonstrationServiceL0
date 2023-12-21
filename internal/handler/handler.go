package handler

import (
	"DemonstrationServiceL0/internal/caching"
	"DemonstrationServiceL0/internal/service"
	"database/sql"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var id int64
var mu sync.Mutex

func Handler(msg *stan.Msg, db *sql.DB) {
	mu.Lock()
	currentID := atomic.AddInt64(&id, 1)

	order := service.DecodeJSON(msg.Data)
	log.Printf("Received message: %s\n", order.OrderUID)
	if err := insertData(db, order); err != nil {
		log.Println("Error inserting data into database:", err)
	}

	orderCopy := *order
	caching.SetCache(int(currentID), orderCopy, 10*time.Minute)
	mu.Unlock()
}

// Вставка данных в базу данных
func insertData(db *sql.DB, order *service.Order) error {
	deliveryUid := uuid.NewString()
	paymentUid := uuid.NewString()
	itemsUid := uuid.NewString()

	// Вставка данных в таблице orders
	_, err := db.Exec(insertInOrders,
		order.OrderUID, order.TrackNumber, order.Entry, deliveryUid, paymentUid, itemsUid, order.Locale,
		order.InternalSig, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID,
		order.DateCreated, order.OofShard,
	)
	if err != nil {
		log.Println(1)
	}

	// Вставка данных в таблице delivery
	_, err = db.Exec(insertInDelivery,
		deliveryUid, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email,
	)
	if err != nil {
		return err
	}

	// Вставка данных в таблице payment
	_, err = db.Exec(insertInPayment,
		paymentUid, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency,
		order.Payment.Provider, order.Payment.Amount, time.Unix(order.Payment.PaymentDt, 0),
		order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee,
	)
	if err != nil {
		return err
	}

	// Вставка данных в таблице item
	for _, item := range order.Items {
		_, err = db.Exec(insertInItems,
			itemsUid, item.ChrtID, item.TrackNumber, item.Price, item.RID,
			item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID,
			item.Brand, item.Status,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
