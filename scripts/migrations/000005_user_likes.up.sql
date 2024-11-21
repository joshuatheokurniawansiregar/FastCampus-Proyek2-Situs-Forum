CREATE TABLE IF NOT EXISTS user_likes(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    user_id BIGINT NOT NULL,
    post_id BIGINT NOT NULL,
    does_like BOOLEAN NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    CONSTRAINT fk_user_id_user_likes FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_post_id_user_likes FOREIGN KEY(post_id) REFERENCES posts(id)    
)DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;