package main

import (
	"os"

	"github.com/go-chi/chi/v5"

	app "github.com/diwise/api-notify/internal/pkg/application"
	repo "github.com/diwise/api-notify/internal/pkg/database"
)

func main() {
	r := chi.NewRouter()
	database := repo.NewDatabase(os.Getenv("DATABASE_URL"))
	a := app.NewApplication(r, *database)
	a.Start(os.Getenv("SERVICE_PORT"))
}
