CREATE TABLE IF NOT EXISTS _user_addresses (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    city VARCHAR(255) NOT NULL,
    province VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    license_plate VARCHAR(255) NOT NULL,
    description TEXT,
    user_id INT REFERENCES _users(id)
)