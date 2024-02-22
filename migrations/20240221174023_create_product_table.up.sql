CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    purchase_price NUMERIC(10,2) NOT NULL CHECK(purchase_price > 0),
    sales_price NUMERIC(10,2) NOT NULL CHECK(purchase_price > 0),
    stock INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);