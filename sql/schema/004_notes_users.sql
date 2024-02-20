-- +goose Up
ALTER TABLE notes
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE notes
DROP CONSTRAINT IF EXISTS fk_user_id;
