package database

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
}

func (u *User) GetUserDTO() UserDTO {
	return UserDTO{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Username:  u.Username,
	}
}

type NoteDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (n *Note) GetNoteDTO() NoteDTO {
	return NoteDTO{
		ID:        n.ID,
		Title:     n.Title,
		Body:      n.Body,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}
