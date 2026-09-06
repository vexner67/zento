-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    status TEXT NOT NULL,
    currency_code CHAR(3) NOT NULL,
    total_units BIGINT NOT NULL,
    total_nanos INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price_units BIGINT NOT NULL,
    unit_price_nanos INTEGER NOT NULL,

    PRIMARY KEY (order_id, product_id)
);

-- +goose Down
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_items;
