-- +goose Up
CREATE TABLE notes (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    body VARCHAR NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE notes;