package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EwokOwie/dog-api/internal/models"
)

type application struct {
	logger  *slog.Logger
	animals *models.AnimalService
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	animals := models.NewAnimalService()

	app := &application{
		logger:  logger,
		animals: animals,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Channel to listen for errors from the server
	shutdownError := make(chan error)

	// Start a goroutine to listen for shutdown signals
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	logger.Info("starting server", "addr", *addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err.Error())
		os.Exit(1)
	}

	err = <-shutdownError
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("stopped server", "addr", *addr)
}
