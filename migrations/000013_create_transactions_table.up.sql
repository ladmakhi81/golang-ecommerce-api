CREATE TABLE IF NOT EXISTS _transactions (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT NOT NULL REFERENCES _users(id),
    payment_id INT NOT NULL REFERENCES _payments(id),
    order_id INT NOT NULL REFERENCES _orders(id),
    authority VARCHAR(255) NOT NULL,
    ref_id INT NOT NULL,
    amount DECIMAL NOT NULL
)