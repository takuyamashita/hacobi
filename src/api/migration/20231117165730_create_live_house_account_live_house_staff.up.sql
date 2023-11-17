CREATE TABLE IF NOT EXISTS live_house_account_live_house_staff (
    live_house_account_id VARCHAR(36) NOT NULL,
    live_house_staff_id VARCHAR(36) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    INDEX live_house_account_id (live_house_account_id),
    UNIQUE KEY live_house_account_id_live_house_staff_id (live_house_account_id, live_house_staff_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;