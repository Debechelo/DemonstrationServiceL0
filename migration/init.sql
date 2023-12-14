CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE orders
(
    order_uid          UUID PRIMARY KEY,
    track_number       VARCHAR(255) UNIQUE,
    entry              VARCHAR(255),
    delivery_uid       UUID,
    payment            UUID,
    items              UUID,
    locale             VARCHAR(12),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shard_key          VARCHAR(255),
    sm_id              NUMERIC,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR(255)
);

CREATE TABLE delivery
(
    delivery_uid UUID PRIMARY KEY,
    name         VARCHAR(255),
    phone        VARCHAR(20),
    zip          VARCHAR(20),
    city         VARCHAR(255),
    address      VARCHAR(255),
    region       VARCHAR(255),
    email        VARCHAR(255),
    FOREIGN KEY (delivery_uid) REFERENCES orders(delivery_uid)
);

CREATE TABLE payment
(
    payment_uid   UUID PRIMARY KEY,
    transaction   VARCHAR(255),
    request_id    VARCHAR(255),
    currency      VARCHAR(5),
    provider      VARCHAR(255),
    amount        INT,
    payment_dt    TIMESTAMP,
    bank          VARCHAR(255),
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT,
    FOREIGN KEY (payment_uid) REFERENCES orders (payment)
);

CREATE TABLE items
(
    item_uid     UUID PRIMARY KEY,
    order_uid    UUID,
    chrt_id      INT,
    track_number VARCHAR(255),
    price        INT,
    rid          VARCHAR(255),
    item_name    VARCHAR(255),
    sale         INT,
    size         VARCHAR(10),
    total_price  INT,
    nm_id        INT,
    brand        VARCHAR(255),
    status       INT,
    FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);