CREATE TABLE IF NOT EXISTS refresh_tokens(
    id BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    refresh_token TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    crated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    CONSTRAINT fk_user_id_refresh_tokens FOREIGN KEY(user_id) REFERENCES users(id)
)DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;