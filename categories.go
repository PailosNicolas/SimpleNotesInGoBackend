package main

import (
	"encoding/json"
	"net/http"

	"github.com/PailosNicolas/SimpleNotesInGoBackend/helpers"
	"github.com/PailosNicolas/SimpleNotesInGoBackend/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.Name == "" {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Missing parameters.")
		return
	}

	category, err := cfg.DB.CreateCategory(r.Context(), database.CreateCategoryParams{
		ID:     uuid.New(),
		Name:   params.Name,
		UserID: user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating category")
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, category.GetDTO())
}

func (cfg *apiConfig) HandlerGetCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := helpers.PaginationParams{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	categories, err := cfg.DB.GetCategoriesByUser(r.Context(), user.ID)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting categories")
		return
	}

	paginatedNotes := helpers.PaginateResult(database.GetCategorySliceDTOs(categories), helpers.PaginationParams{
		Page:     params.Page,
		PageSize: params.PageSize,
	})

	helpers.RespondWithJSON(w, http.StatusOK, paginatedNotes)
}

func (cfg *apiConfig) HandlerDeleteCategory(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		CategoryID string `json:"category_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	categoryUuid, err := uuid.Parse(params.CategoryID)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	err = cfg.DB.DeleteCategoryById(r.Context(), database.DeleteCategoryByIdParams{
		ID:     categoryUuid,
		UserID: user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error deleting categories")
		return
	}

	helpers.RespondWithOK(w)
}

func (cfg *apiConfig) HandlerAssingCategoryToNote(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		CategoryID string `json:"category_id"`
		NoteID     string `json:"note_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.CategoryID == "" || params.NoteID == "" {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Missing parameters.")
		return
	}

	categoryUuid, err := uuid.Parse(params.CategoryID)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	NoteUuid, err := uuid.Parse(params.NoteID)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	category, err := cfg.DB.GetCategoryById(r.Context(), database.GetCategoryByIdParams{
		ID:     categoryUuid,
		UserID: user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error assigning category")
		return
	}

	note, err := cfg.DB.GetNoteById(r.Context(), database.GetNoteByIdParams{
		ID:     NoteUuid,
		UserID: user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error assigning category")
		return
	}

	_, err = cfg.DB.CreateNoteCategory(r.Context(), database.CreateNoteCategoryParams{
		NoteID:     note.Note.ID,
		CategoryID: category.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error assigning category")
		return
	}

	helpers.RespondWithOK(w)
}
