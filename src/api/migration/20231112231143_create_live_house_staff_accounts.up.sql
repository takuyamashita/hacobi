CREATE TABLE IF NOT EXISTS live_house_staff_accounts (
    id VARCHAR(36) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);