CREATE TABLE IF NOT EXISTS live_house_staff_profiles (
    live_house_staff_account_id VARCHAR(36) NOT NULL,
    display_name VARCHAR(15) NOT NULL,

    INDEX (live_house_staff_account_id),
    FOREIGN KEY (live_house_staff_account_id) REFERENCES live_house_staff_accounts(id) ON DELETE CASCADE,

    UNIQUE (live_house_staff_account_id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;