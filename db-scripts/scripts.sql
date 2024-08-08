CREATE table product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(50) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
)

INSERT INTO product(product_name, price) VALUES('Sushi', 100)