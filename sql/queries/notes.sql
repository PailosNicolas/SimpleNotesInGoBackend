-- name: CreateNote :one
INSERT INTO notes (id, title, body, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateNoteTitleBody :one
UPDATE notes
SET title = $1, body = $2, updated_at = $3
WHERE id = $4
RETURNING *;

-- name: GetNoteById :one
SELECT *
FROM notes
WHERE id=$1 AND user_id=$2;