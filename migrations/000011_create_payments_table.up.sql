CREATE TYPE payment_status AS ENUM ('Pending', 'Success', 'Failed');

CREATE TABLE IF NOT EXISTS _payments (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_id INT NOT NULL REFERENCES _users(id),
    status payment_status NOT NULL,
    status_changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    order_id INT NOT NULL REFERENCES _orders(id),
    amount DECIMAL NOT NULL,
    authority VARCHAR(255) NOT NULL
)