CREATE TYPE user_role AS ENUM ('Customer', 'Admin', 'Vendor');

CREATE TABLE IF NOT EXISTS _users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_role user_role,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    national_id VARCHAR(255),
    postal_code VARCHAR(255),
    address TEXT,
    is_verified BOOLEAN DEFAULT FALSE,
    verified_by_id INT REFERENCES "_users"("id")
)