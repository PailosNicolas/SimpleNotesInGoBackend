package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PailosNicolas/SimpleNotesInGoBackend/helpers"
	"github.com/PailosNicolas/SimpleNotesInGoBackend/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateNote(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.Title == "" || params.Body == "" {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Missing parameters.")
		return
	}

	note, err := cfg.DB.CreateNote(r.Context(), database.CreateNoteParams{
		Title:     params.Title,
		Body:      params.Body,
		UserID:    user.ID,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating note")
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, note.GetNoteDTO())
}

func (cfg *apiConfig) HandlerUpdateNote(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
		Body  string    `json:"body"`
	}
	var emptyUUID uuid.UUID

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.ID == emptyUUID {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Missing id.")
		return
	}

	note, err := cfg.DB.GetNoteById(r.Context(), database.GetNoteByIdParams{
		ID:     params.ID,
		UserID: user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting note")
		return
	}

	if params.Body == "" {
		params.Body = note.Body
	}

	if params.Title == "" {
		params.Title = note.Title
	}

	updatedNote, err := cfg.DB.UpdateNoteTitleBody(r.Context(), database.UpdateNoteTitleBodyParams{
		Title:     params.Title,
		Body:      params.Body,
		UpdatedAt: time.Now(),
		ID:        note.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error updating note")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, updatedNote.GetNoteDTO())
}
