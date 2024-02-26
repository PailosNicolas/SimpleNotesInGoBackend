// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: categories.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (id, name, user_id)
VALUES ($1, $2, $3)
RETURNING id, name, user_id
`

type CreateCategoryParams struct {
	ID     uuid.UUID
	Name   string
	UserID uuid.UUID
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.ID, arg.Name, arg.UserID)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.UserID)
	return i, err
}

const deleteCategoryById = `-- name: DeleteCategoryById :exec
DELETE FROM categories
WHERE id = $1 AND user_id = $2
`

type DeleteCategoryByIdParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteCategoryById(ctx context.Context, arg DeleteCategoryByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteCategoryById, arg.ID, arg.UserID)
	return err
}

const getCategoriesByUser = `-- name: GetCategoriesByUser :many
SELECT id, name, user_id
FROM categories
WHERE user_id=$1
ORDER BY name DESC
`

func (q *Queries) GetCategoriesByUser(ctx context.Context, userID uuid.UUID) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name, user_id
FROM categories
WHERE id=$1 AND user_id=$2
`

type GetCategoryByIdParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) GetCategoryById(ctx context.Context, arg GetCategoryByIdParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryById, arg.ID, arg.UserID)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.UserID)
	return i, err
}
