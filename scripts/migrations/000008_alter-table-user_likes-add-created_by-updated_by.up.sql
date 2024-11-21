ALTER TABLE user_likes
ADD COLUMN created_by LONGTEXT NULL AFTER updated_at,
ADD COLUMN updated_by LONGTEXT NULL AFTER created_by;