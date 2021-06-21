package main

import (
	"log"
	"os"

	"github.com/go-chi/chi/v5"

	app "github.com/diwise/api-notify/internal/pkg/application"
	repo "github.com/diwise/api-notify/internal/pkg/database"

	mq "github.com/diwise/messaging-golang/pkg/messaging"
)

func main() {
	r := chi.NewRouter()
	database := repo.NewDatabase(os.Getenv("DATABASE_URL"))

	ctx, _ := mq.Initialize(mq.Config{
		Host:        os.Getenv("RABBITMQ_HOST"),
		User:        os.Getenv("RABBITMQ_USER"),
		Password:    os.Getenv("RABBITMQ_PASSWORD"),
		ServiceName: "api-notify",
	})

	defer ctx.Close()

	a := app.NewApplication(r, database, ctx)
	log.Fatal(a.Start(os.Getenv("SERVICE_PORT")))
}
