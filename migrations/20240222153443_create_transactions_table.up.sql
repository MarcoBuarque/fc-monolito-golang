CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    status VARCHAR(10) NOT NULL CHECK(document_type IN ('pending', 'approved', 'declined')),
    amount NUMERIC(10,2) NOT NULL CHECK(amount > 0),

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);