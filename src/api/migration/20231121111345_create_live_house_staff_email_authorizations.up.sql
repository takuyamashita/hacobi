CREATE TABLE IF NOT EXISTS live_house_staff_email_authorizations (
    email VARCHAR(255) NOT NULL,
    token VARCHAR(128) NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);