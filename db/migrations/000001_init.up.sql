CREATE TABLE "order"
(
    id         SERIAL PRIMARY KEY,
    order_uid  TEXT UNIQUE NOT NULL,
    data       JSONB,
    created_at TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW())
);
