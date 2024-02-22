CREATE TYPE transaction_status AS ENUM ("pending", "approved", "declined");

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    status transaction_status NOT NULL,
    amount NUMERIC(10,2) NOT NULL CHECK(amount > 0),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);