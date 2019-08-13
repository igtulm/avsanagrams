package avsanagrams

import (
	"fmt"
	"log"
	"net/http"

	"avsanagrams/internal/app"
	"avsanagrams/internal/app/avsanagrams/api"
	"avsanagrams/internal/app/avsanagrams/config"
	"avsanagrams/internal/pkg/storage"
)

type application struct {
	cfg *config.Cfg
}

func New(cfg *config.Cfg) app.App {
	return &application{
		cfg: cfg,
	}
}

func (a *application) serverAddress() string {
	return fmt.Sprintf("%s:%s", a.cfg.Host, a.cfg.Port)
}

func (a *application) registerRoutes() {
	store := storage.New()
	handlers := api.New(store)

	http.HandleFunc("/load", handlers.Load)
	http.HandleFunc("/get", handlers.Get)
}

func (a *application) Start() error {
	a.registerRoutes()

	addr := a.serverAddress()

	log.Printf("The server is starting on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}
