package main

import (
	"log"
	"github.com/hmprgm/swocial/internal/env"
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
		addr: env.GetString("ADDR",":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}