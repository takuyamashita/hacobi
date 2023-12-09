CREATE TABLE IF NOT EXISTS live_house_staff_account_credential_relations (
    live_house_staff_account_id VARCHAR(36) NOT NULL,
    public_key_id VARCHAR(128) NOT NULL,

    INDEX (live_house_staff_account_id),
    FOREIGN KEY (live_house_staff_account_id) REFERENCES live_house_staff_accounts (id) ON DELETE CASCADE,

    INDEX (public_key_id),
    FOREIGN KEY (public_key_id) REFERENCES account_credentials (public_key_id) ON DELETE CASCADE,

    UNIQUE (live_house_staff_account_id, public_key_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
