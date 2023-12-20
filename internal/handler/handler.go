package handler

import (
	"DemonstrationServiceL0/internal/service"
	"database/sql"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func Handler(msg *stan.Msg, db *sql.DB) {
	order := service.DecodeJSON(msg.Data)
	log.Printf("Received message: %s\n", order.OrderUID)
	if err := insertData(db, order); err != nil {
		log.Println("Error inserting data into database:", err)
	}
}

// Вставка данных в базу данных
func insertData(db *sql.DB, order *service.Order) error {
	deliveryUid := uuid.NewString()
	paymentUid := uuid.NewString()
	itemsUid := uuid.NewString()

	// Вставка данных в таблице orders
	_, err := db.Exec(`
		INSERT INTO orders (
			order_uid, track_number, entry, delivery_uid, payment_uid, item_uid, locale, internal_signature,
			customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
		)`,
		order.OrderUID, order.TrackNumber, order.Entry, deliveryUid, paymentUid, itemsUid, order.Locale,
		order.InternalSig, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID,
		order.DateCreated, order.OofShard,
	)
	if err != nil {
		log.Println(1)
		return err
	}

	// Вставка данных в таблице delivery
	_, err = db.Exec(`
		INSERT INTO delivery (
			delivery_uid, name, phone, zip, city, address, region, email
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)`,
		deliveryUid, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email,
	)
	if err != nil {
		log.Println(2)
		return err
	}

	// Вставка данных в таблице payment
	_, err = db.Exec(`
		INSERT INTO payment (
			payment_uid, transaction, request_id, currency, provider, amount, payment_dt, bank,
			delivery_cost, goods_total, custom_fee
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)`,
		paymentUid, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency,
		order.Payment.Provider, order.Payment.Amount, time.Unix(order.Payment.PaymentDt, 0),
		order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee,
	)
	if err != nil {
		log.Println(3)
		return err
	}

	// Вставка данных в таблице item
	_, err = db.Exec(`
			INSERT INTO items (
				item_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
			) VALUES (
				 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
			)`,
		itemsUid, order.Item.ChrtID, order.Item.TrackNumber, order.Item.Price, order.Item.RID,
		order.Item.Name, order.Item.Sale, order.Item.Size, order.Item.TotalPrice, order.Item.NmID,
		order.Item.Brand, order.Item.Status,
	)
	if err != nil {
		log.Println(4)
		return err
	}

	return nil
}
