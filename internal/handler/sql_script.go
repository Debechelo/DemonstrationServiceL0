package handler

const (
	insertInOrders = `INSERT INTO orders (
				order_uid, track_number, entry, delivery_uid, payment_uid, item_uid, locale, internal_signature,
				customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
		)`
	insertInDelivery = `INSERT INTO delivery (
				delivery_uid, name, phone, zip, city, address, region, email
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8
		)`
	insertInPayment = `INSERT INTO payment (
				payment_uid, transaction, request_id, currency, provider, amount, payment_dt, bank,
				delivery_cost, goods_total, custom_fee
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)`
	insertInItems = `INSERT INTO items (
				item_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)`
)
