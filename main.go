package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/yusuffranklin/notes-api/database"
	"github.com/yusuffranklin/notes-api/logger"
	"github.com/yusuffranklin/notes-api/routes"
	"go.uber.org/zap"
)

func main() {
	// Connect to db
	database.InitDB()

	// Create table if not exists
	// database.CreateTables()

	router := mux.NewRouter()
	routes.RegisteredRoutes(router)

	logger.Info("Server Running", zap.String("url", "http://localhost:8080"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
