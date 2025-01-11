CREATE TABLE IF NOT EXISTS _categories (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) UNIQUE,
    icon TEXT NOT NULL,
    parent_category_id INT REFERENCES "_categories"("id")
);