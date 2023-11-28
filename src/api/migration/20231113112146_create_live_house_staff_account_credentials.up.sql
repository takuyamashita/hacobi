CREATE TABLE IF NOT EXISTS live_house_staff_account_credentials (
    live_house_staff_account_id VARCHAR(36) NOT NULL,
    public_key_id VARCHAR(128) NOT NULL,
    public_key VARCHAR(128) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);