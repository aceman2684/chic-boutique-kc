CREATE TABLE IF NOT EXISTS "address" (
    address_id SERIAL PRIMARY KEY,
    line_one VARCHAR(255) NOT NULL,
    line_two VARCHAR(255),
    city VARCHAR(255) NOT NULL,
    "state" VARCHAR(255) NOT NULL,
    zip_code VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS "user" (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS customer (
    customer_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL UNIQUE,
    FOREIGN KEY (user_id)
        REFERENCES "user" (user_id)
);

CREATE TABLE IF NOT EXISTS customer_address (
    address_id INT NOT NULL,
    customer_id INT NOT NULL,
    PRIMARY KEY (address_id, customer_id),
    FOREIGN KEY (address_id)
        REFERENCES "address" (address_id),
    FOREIGN KEY (customer_id)
        REFERENCES customer (customer_id)
);

CREATE TABLE IF NOT EXISTS product (
    product_id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS product_price (
    product_price_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    dollars INT NOT NULL,
    cents INT NOT NULL,
    active_ind BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (product_id)
        REFERENCES product (product_id)
);

CREATE TABLE IF NOT EXISTS "order" (
    order_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    parent_order_id INT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (customer_id)
        REFERENCES customer (customer_id),
    FOREIGN KEY (parent_order_id)
        REFERENCES "order" (order_id)
);

CREATE TABLE IF NOT EXISTS order_product (
    order_id INT NOT NULL,
    product_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS order_price (
    order_id INT NOT NULL,
    dollars INT NOT NULL,
    cents INT NOT NULL,
    active_ind BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (order_id)
        REFERENCES "order" (order_id)
);