package main

import (
	"net/http"
	"os"
	"time"

	v1 "github.com/stianfro/chi-no-wadachi/api/v1"
	utils "github.com/stianfro/chi-no-wadachi/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	utils.SetEnv()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/v1/healthz", v1.HealthZ)

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to start server")
	}
}
