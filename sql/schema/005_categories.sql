-- +goose Up
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE note_categories (
    note_id UUID REFERENCES notes(id) ON DELETE CASCADE,
    category_id UUID REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (note_id, category_id)
);

-- +goose Down
DROP TABLE IF EXISTS note_categories;
DROP TABLE IF EXISTS categories;
