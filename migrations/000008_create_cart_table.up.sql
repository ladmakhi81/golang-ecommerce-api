CREATE TABLE IF NOT EXISTS _carts(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_id INT NOT NULL REFERENCES _products(id),
    customer_id INT NOT NULL REFERENCES _users(id),
    price_item_id INT NOT NULL REFERENCES _product_prices(id),
    quantity INT NOT NULL
)