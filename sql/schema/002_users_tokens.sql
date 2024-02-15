-- +goose Up
ALTER TABLE users
ADD COLUMN token VARCHAR,
ADD COLUMN refresh_token VARCHAR;

-- +goose Down
ALTER TABLE users
DROP COLUMN token,
DROP COLUMN refresh_token;
