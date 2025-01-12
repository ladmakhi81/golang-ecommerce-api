CREATE TABLE IF NOT EXISTS _products(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    base_price DECIMAL NOT NULL,
    tags TEXT [],
    is_confirmed BOOLEAN DEFAULT FALSE,
    confirmed_at TIMESTAMP,
    confirmed_by_id INT REFERENCES "_users"("id"),
    category_id INT REFERENCES "_categories"("id"),
    vendor_id INT REFERENCES "_users"("id")
)