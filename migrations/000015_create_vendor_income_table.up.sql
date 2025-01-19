CREATE TABLE IF NOT EXISTS _vendor_incomes (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_id INT NOT NULL REFERENCES _users(id),
    order_amount DECIMAL NOT NULL,
    fee_amount DECIMAL NOT NULL,
    income_amount DECIMAL NOT NULL,
    order_item INT REFERENCES _order_items(id),
    transaction_id INT REFERENCES _transactions(id)
)