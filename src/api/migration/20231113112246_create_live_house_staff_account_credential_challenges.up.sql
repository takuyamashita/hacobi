CREATE TABLE IF NOT EXISTS live_house_staff_account_credential_challenges (
    live_house_staff_account_id VARCHAR(36) NOT NULL,
    challenge VARCHAR(128) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);