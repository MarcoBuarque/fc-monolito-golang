CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    client_id INTEGER REFERENCES clients(id),

    status VARCHAR(10) NOT NULL CHECK(document_type IN ('pending', 'approved', 'declined')),
    total DECIMAL(10, 2) NOT NULL,
   
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);


-- considering that the sales price may change over time
CREATE TABLE IF NOT EXISTS order_products (
    order_id VARCHAR(255) REFERENCES orders(id),
    product_id VARCHAR(255) REFERENCES products(id),
    quantity INTEGER NOT NULL CHECK(quantity > 0),
    price DECIMAL(10, 2) NOT NULL CHECK(price > 0),
    PRIMARY KEY (order_id, product_id)
);
