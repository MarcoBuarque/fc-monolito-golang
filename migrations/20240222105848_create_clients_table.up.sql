CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
-- Document
    document_type VARCHAR(10) NOT NULL CHECK(document_type IN ('RG', 'CPF', 'CNPJ')),
    document_number VARCHAR(20) NOT NULL UNIQUE,

    -- Address
    email VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    zipCode VARCHAR(10) NOT NULL,


    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);