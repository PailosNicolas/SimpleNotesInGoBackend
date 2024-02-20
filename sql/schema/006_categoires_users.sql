-- +goose Up
ALTER TABLE categories
ADD COLUMN user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE categories
DROP COLUMN user_id;
