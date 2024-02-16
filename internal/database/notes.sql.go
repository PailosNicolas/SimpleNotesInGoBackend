// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: notes.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (id, title, body, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, title, body, user_id, created_at, updated_at
`

type CreateNoteParams struct {
	ID        uuid.UUID
	Title     string
	Body      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, createNote,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNoteById = `-- name: GetNoteById :one
SELECT id, title, body, user_id, created_at, updated_at
FROM notes
WHERE id=$1 AND user_id=$2
`

type GetNoteByIdParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) GetNoteById(ctx context.Context, arg GetNoteByIdParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, getNoteById, arg.ID, arg.UserID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateNoteTitleBody = `-- name: UpdateNoteTitleBody :one
UPDATE notes
SET title = $1, body = $2, updated_at = $3
WHERE id = $4
RETURNING id, title, body, user_id, created_at, updated_at
`

type UpdateNoteTitleBodyParams struct {
	Title     string
	Body      string
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateNoteTitleBody(ctx context.Context, arg UpdateNoteTitleBodyParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, updateNoteTitleBody,
		arg.Title,
		arg.Body,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
