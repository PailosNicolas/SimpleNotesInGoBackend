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
SELECT sqlc.embed(n), json_agg(c) as categories
FROM notes n
LEFT JOIN note_categories nc ON nc.note_id = n.id
LEFT JOIN categories c ON c.id = nc.category_id
WHERE n.id = $1 AND n.user_id = $2
GROUP BY n.id;

-- name: GetNotesByUser :many
SELECT
  sqlc.embed(n),
  json_agg(c) as categories
FROM
  notes n
LEFT JOIN
  note_categories nc ON nc.note_id = n.id
LEFT JOIN
  categories c ON c.id = nc.category_id
WHERE
  n.user_id = $1
  AND ($2 IS FALSE OR nc.category_id = ANY($3::uuid[]))
GROUP BY
  n.id
ORDER BY
  n.created_at DESC;
