CREATE TABLE IF NOT EXISTS account_credentials (
    public_key_id VARCHAR(128) NOT NULL,
    public_key VARCHAR(128) NOT NULL,
    attestation_type VARCHAR(128) NOT NULL,
    transport INT(32) UNSIGNED NOT NULL,

    user_present TINYINT(1) NOT NULL,
    user_verified TINYINT(1) NOT NULL,
    backup_eligible TINYINT(1) NOT NULL,
    backup_state TINYINT(1) NOT NULL,

    aaguid VARCHAR(128) NOT NULL,
    sign_count INT(32) UNSIGNED NOT NULL,
    attachment VARCHAR(128) NOT NULL,

    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (public_key_id)
);
