CREATE TABLE delivery
(
    id      SERIAL PRIMARY KEY,
    name    TEXT,
    phone   TEXT,
    zip     TEXT,
    city    TEXT,
    address TEXT,
    region  TEXT,
    email   TEXT UNIQUE NOT NULL
);

CREATE TABLE payment
(
    id            SERIAL PRIMARY KEY,
    transaction   TEXT UNIQUE NOT NULL,
    request_id    TEXT,
    currency      TEXT,
    provider      TEXT,
    amount        DECIMAL(10, 2),
    payment_dt    INT,
    bank          TEXT,
    delivery_cost DECIMAL(10, 2),
    goods_total   DECIMAL(10, 2),
    custom_fee    DECIMAL(10, 2)
);

CREATE TABLE item
(
    id           SERIAL PRIMARY KEY,
    chrt_id      INT UNIQUE NOT NULL,
    track_number TEXT,
    price        DECIMAL(10, 2),
    rid          TEXT,
    item_name    TEXT,
    sale         INT,
    size         TEXT,
    total_price  DECIMAL(10, 2),
    nm_id        INT,
    brand        TEXT,
    status       INT
);

CREATE TABLE customer
(
    customer_uid TEXT PRIMARY KEY,
    email        TEXT UNIQUE NOT NULL
);

CREATE TABLE "order"
(
    id                 SERIAL PRIMARY KEY,
    order_uid          TEXT,
    track_number       TEXT,
    entry              TEXT,
    delivery_id        INT  REFERENCES delivery (id) ON DELETE SET NULL,
    payment_id         INT  REFERENCES payment (id) ON DELETE SET NULL,
    locale             TEXT,
    internal_signature TEXT,
    customer_id        TEXT NOT NULL REFERENCES customer (customer_uid) ON DELETE CASCADE,
    delivery_service   TEXT,
    shard_key          TEXT,
    sm_id              INT,
    date_created       TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW()),
    updated_at         TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW()),
    oof_shard          TEXT
);

CREATE INDEX ON "order" (delivery_id);
CREATE INDEX ON "order" (payment_id);
CREATE INDEX ON "order" (customer_id);

CREATE TABLE order_item
(
    order_id INT NOT NULL REFERENCES "order" (id) ON DELETE CASCADE,
    item_id  INT NOT NULL REFERENCES item (id) ON DELETE CASCADE,
    PRIMARY KEY (order_id, item_id)
);

CREATE
    OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = timezone('Europe/Moscow'::text, NOW());
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER user_updated_at_trigger
    BEFORE UPDATE
    ON "order"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();