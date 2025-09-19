package routes

import (
	"github.com/gorilla/mux"
	"github.com/yusuffranklin/notes-api/handlers"
)

func RegisteredRoutes(router *mux.Router) {
	router.HandleFunc("/notes", handlers.CreateNoteHandler).Methods("POST")
	router.HandleFunc("/notes/{id}", handlers.GetNoteHandler).Methods("GET")
	router.HandleFunc("/notes/{id}", handlers.UpdateNoteHandler).Methods("PUT")
	router.HandleFunc("/notes/{id}", handlers.DeleteNoteHandler).Methods("DELETE")
}
