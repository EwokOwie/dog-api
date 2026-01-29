package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

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

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
