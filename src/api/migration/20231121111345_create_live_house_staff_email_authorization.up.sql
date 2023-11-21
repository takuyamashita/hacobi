CREATE TABLE IF NOT EXISTS live_house_staff_email_authorization (
    email VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);