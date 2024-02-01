CREATE TABLE delivery
(
    id      SERIAL PRIMARY KEY,
    name    TEXT,
    phone   TEXT,
    zip     TEXT,
    city    TEXT,
    address TEXT,
    region  TEXT,
    email   TEXT
);

CREATE TABLE payment
(
    id            SERIAL PRIMARY KEY,
    transaction   TEXT,
    request_id    TEXT,
    currency      TEXT,
    provider      TEXT,
    amount        DECIMAL(10, 2),
    payment_dt    TIMESTAMP,
    bank          TEXT,
    delivery_cost DECIMAL(10, 2),
    goods_total   DECIMAL(10, 2),
    custom_fee    DECIMAL(10, 2)
);

CREATE TABLE item
(
    id           SERIAL PRIMARY KEY,
    chrt_id      INTEGER,
    track_number TEXT,
    price        DECIMAL(10, 2),
    rid          TEXT,
    item_name    TEXT,
    sale         INTEGER,
    size         VARCHAR(50),
    total_price  DECIMAL(10, 2),
    nm_id        INTEGER,
    brand        TEXT,
    status       INTEGER
);

CREATE TABLE "order"
(
    id                 SERIAL PRIMARY KEY,
    order_uid          TEXT,
    track_number       TEXT,
    entry              TEXT,
    delivery_id        INTEGER REFERENCES delivery (delivery_id),
    payment_id         INTEGER REFERENCES payment (payment_id),
    locale             TEXT,
    internal_signature TEXT,
    customer_id        TEXT,
    delivery_service   TEXT,
    shard_key          TEXT,
    sm_id              INTEGER,
    date_created       TIMESTAMP,
    oof_shard          TEXT
);

CREATE TABLE order_item
(
    order_id INTEGER REFERENCES "order" (id),
    item_id  INTEGER REFERENCES item (id),
    PRIMARY KEY (order_id, item_id)
);
