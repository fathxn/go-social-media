package main

import (
	"go-social-media/internal/env"
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	os.LookupEnv("")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
