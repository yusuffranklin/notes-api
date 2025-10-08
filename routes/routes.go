package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yusuffranklin/notes-api/handlers"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func RegisteredRoutes(router *mux.Router) {
	router.Handle("/notes", otelhttp.NewHandler(http.HandlerFunc(handlers.CreateNoteHandler), "CreateNote")).Methods("POST")
	router.Handle("/notes/{id}", otelhttp.NewHandler(http.HandlerFunc(handlers.GetNoteHandler), "GetNote")).Methods("GET")
	router.Handle("/notes/{id}", otelhttp.NewHandler(http.HandlerFunc(handlers.UpdateNoteHandler), "UpdateNote")).Methods("PUT")
	router.Handle("/notes/{id}", otelhttp.NewHandler(http.HandlerFunc(handlers.DeleteNoteHandler), "DeleteNote")).Methods("DELETE")
	// router.HandleFunc("/notes", handlers.CreateNoteHandler).Methods("POST")
	// router.HandleFunc("/notes/{id}", handlers.GetNoteHandler).Methods("GET")
	// router.HandleFunc("/notes/{id}", handlers.UpdateNoteHandler).Methods("PUT")
	// router.HandleFunc("/notes/{id}", handlers.DeleteNoteHandler).Methods("DELETE")
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")
}
