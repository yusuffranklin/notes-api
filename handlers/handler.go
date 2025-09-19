package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yusuffranklin/notes-api/database"
	"github.com/yusuffranklin/notes-api/logger"
	"github.com/yusuffranklin/notes-api/models"
	"go.uber.org/zap"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming request", zap.String("handler", "CreateNoteHandler"))

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		logger.Warn("Failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.Db.QueryRow(
		"INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id",
		note.Title, note.Content,
	).Scan(&note.ID)
	if err != nil {
		logger.Error("DB insert failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Note created", zap.Int("id", note.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming request", zap.String("handler", "getNote"))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Warn("Invalid note ID", zap.Error(err))
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note models.Note
	err = database.Db.QueryRow("SELECT id, title, content FROM notes WHERE id=$1", id).
		Scan(&note.ID, &note.Title, &note.Content)

	if err == sql.ErrNoRows {
		logger.Error("Note not found", zap.Error(err))
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	} else if err != nil {
		logger.Error("DB query failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Note retrieved", zap.Any("note", note))
	json.NewEncoder(w).Encode(note)
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming requeest", zap.String("handler", "updateNote"))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Warn("Invalid note ID", zap.Error(err))
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		logger.Error("Failed to decode request", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := database.Db.Exec("UPDATE notes SET title=$1, content=$2 WHERE id=$3",
		updatedNote.Title, updatedNote.Content, id)
	if err != nil {
		logger.Error("Failed to update note", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("Note not found", zap.Error(err))
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	updatedNote.ID = id

	logger.Info("Note updated", zap.Any("note", updatedNote))
	json.NewEncoder(w).Encode(updatedNote)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming request", zap.String("handler", "deleteNote"))

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		logger.Warn("Invalid note ID", zap.Error(err))
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	res, err := database.Db.Exec("DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		logger.Error("Failed to delete note", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("Note not found", zap.Error(err))
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	logger.Info("Note deleted", zap.Int("id", id))
	w.WriteHeader(http.StatusNoContent)
}
