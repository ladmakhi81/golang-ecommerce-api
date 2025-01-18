CREATE TYPE order_status AS ENUM (
    'Pending',
    'Payed',
    'Preparation',
    'Delivery',
    'Done'
);

CREATE TABLE IF NOT EXISTS _orders (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_id INT NOT NULL REFERENCES _users(id),
    status order_status NOT NULL,
    final_price DECIMAL NOT NULL,
    status_changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)