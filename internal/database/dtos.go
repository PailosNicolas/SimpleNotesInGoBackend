package database

import (
	"encoding/json"
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

type NoteWCategoryDTO struct {
	Note     NoteDTO       `json:"note"`
	Category []interface{} `json:"categories"`
}

func (n *GetNotesByUserRow) GetDTO() interface{} {
	var categories []Category
	if err := json.Unmarshal(n.Categories, &categories); err != nil {
		return nil
	}
	return NoteWCategoryDTO{
		Note:     n.Note.GetDTO().(NoteDTO),
		Category: GetCategorySliceDTOs(categories),
	}
}

type CategoryDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
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
		if dto.Name == "" {
			continue
		}
		result = append(result, dto.GetDTO())
	}
	return result
}

func GetNoteWCategorySliceDTO(notes []GetNotesByUserRow) []interface{} {
	var result []interface{}
	for _, dto := range notes {
		result = append(result, dto.GetDTO())
	}
	return result
}
