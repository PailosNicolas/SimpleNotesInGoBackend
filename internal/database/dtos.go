package database

import (
	"time"

	"github.com/google/uuid"
)

type DTO interface {
	GetDTO() interface{}
}

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
}

func (u *User) GetDTO() interface{} {
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

func (n *Note) GetDTO() interface{} {
	return NoteDTO{
		ID:        n.ID,
		Title:     n.Title,
		Body:      n.Body,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}

type CategoryDTO struct {
	ID   uuid.UUID
	Name string
}

func (c *Category) GetDTO() interface{} {
	return CategoryDTO{
		ID:   c.ID,
		Name: c.Name,
	}
}

func GetDTOs(dtos []DTO) []interface{} { // unused for now
	var result []interface{}
	for _, dto := range dtos {
		result = append(result, dto.GetDTO())
	}
	return result
}

func GetNoteSliceDTO(notes []Note) []interface{} {
	var result []interface{}
	for _, dto := range notes {
		result = append(result, dto.GetDTO())
	}
	return result
}

func GetCategorySliceDTOs(categories []Category) []interface{} {
	var result []interface{}
	for _, dto := range categories {
		result = append(result, dto.GetDTO())
	}
	return result
}
