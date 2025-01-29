package main

import (
	"log"

	"github.com/hmprgm/swocial/internal/env"
	"github.com/hmprgm/swocial/internal/store"
	"github.com/joho/godotenv"
)

func setupDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	setupDotEnv()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))
}
