package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"

	app "github.com/diwise/api-notify/internal/pkg/application"
	repo "github.com/diwise/api-notify/internal/pkg/database"

	mq "github.com/diwise/messaging-golang/pkg/messaging"
)

func main() {
	r := chi.NewRouter()

	dbConfig := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=%s&pool_max_conns=10",
		os.Getenv("DIWISE_SQLDB_USER"), os.Getenv("DIWISE_SQLDB_PASS"),
		os.Getenv("DIWISE_SQLDB_HOST"),
		os.Getenv("DIWISE_SQLDB_NAME"),
		os.Getenv("DIWISE_SQLDB_SSLMODE"))
	database, err := repo.NewDatabase(dbConfig)

	if err != nil {
		log.Fatalf("database error: %s", err.Error())
	}

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
