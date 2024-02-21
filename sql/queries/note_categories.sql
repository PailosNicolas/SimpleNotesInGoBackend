-- name: CreateNoteCategory :one
INSERT INTO note_categories (note_id, category_id)
VALUES ($1, $2)
RETURNING *;