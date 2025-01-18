CREATE TABLE IF NOT EXISTS _order_items (
    id BIGSERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES _products(id),
    price_item_id INT NOT NULL REFERENCES _product_prices(id),
    vendor_id INT NOT NULL REFERENCES _users(id),
    customer_id INT NOT NULL REFERENCES _users(id),
    order_id INT NOT NULL REFERENCES _orders(id),
    quantity INT NOT NULL
)