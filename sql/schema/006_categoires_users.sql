-- +goose Up
ALTER TABLE categories
ADD COLUMN user_id UUID REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE categories
DROP COLUMN user_id;
