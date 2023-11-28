CREATE TABLE IF NOT EXISTS live_house_staff_accounts (
    id VARCHAR(36) NOT NULL,
    email VARCHAR(255) NOT NULL,
    is_provisional BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (id),
    UNIQUE INDEX email (email)
);