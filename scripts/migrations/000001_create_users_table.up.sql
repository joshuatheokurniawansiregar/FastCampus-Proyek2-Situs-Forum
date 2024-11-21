CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(500) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL
)DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;