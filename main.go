package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/yusuffranklin/notes-api/database"
	"github.com/yusuffranklin/notes-api/logger"
	"github.com/yusuffranklin/notes-api/opentelemetry"
	"github.com/yusuffranklin/notes-api/routes"
	"go.uber.org/zap"
)

func main() {
	// Connect to db
	database.InitDB()

	if err := run(); err != nil {
		logger.Fatal("Can't run server:", zap.Error(err))
	}
}

func run() error {
	// Handle SIGINT (CTRL+C) gracefully
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry
	otelShutdown, err := opentelemetry.SetupOTelSDK(ctx)
	if err != nil {
		logger.Fatal("Can't set up OpenTelemetry:", zap.Error(err))
	}

	// Handle shutdown properly so nothing leaks
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server
	router := mux.NewRouter()
	routes.RegisteredRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appUrl := fmt.Sprintf("http://localhost:%s", port)

	logger.Info("Server Running", zap.String("url", appUrl))
	log.Fatal(http.ListenAndServe(":8080", router))

	return err
}
