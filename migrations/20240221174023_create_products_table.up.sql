CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    purchase_price NUMERIC(10,2) NOT NULL CHECK(purchase_price > 0),
    sales_price NUMERIC(10,2) NOT NULL CHECK(sales_price > 0),
    stock INT NOT NULL CHECK(check > 0),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);