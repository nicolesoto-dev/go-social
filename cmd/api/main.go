package main

import (
	"log"

	"github.com/nicolesoto-dev/go-social/internal/store"

	"github.com/nicolesoto-dev/go-social/internal/env"
)

func main() {
	cfg := Config{
		addr: env.GetString("ADDR", ":8081"),
	}

	store := store.NewStore(nil)
	app := &application{
		config:  cfg,
		storage: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
