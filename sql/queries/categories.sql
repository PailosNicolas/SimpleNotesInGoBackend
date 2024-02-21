-- name: CreateCategory :one
INSERT INTO categories (id, name, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCategoriesByUser :many
SELECT *
FROM categories
WHERE user_id=$1
ORDER BY name DESC;

-- name: GetCategoryById :one
SELECT *
FROM categories
WHERE id=$1 AND user_id=$2;