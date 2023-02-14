package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	v1 "github.com/stianfro/wotapi/api/v1"
	_ "github.com/stianfro/wotapi/docs"
	"github.com/stianfro/wotapi/models"
	"github.com/stianfro/wotapi/services"
	utils "github.com/stianfro/wotapi/utils"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title WotAPI
// @description  An API that helps you track your hobbies.
// @version 1.0.0
// @license.name MIT
// @BasePath /api/v1
func main() {
	// Environment
	utils.SetEnv()

	// Database
	db, err := utils.InitDB()
	if err != nil {
		log.Err(err).Msg("Failed to connect to database")
	}

	// Interface
	mangastore := &models.MangaStore{DataBase: db}
	service := &services.Service{MangaStore: mangastore}

	// Server
	webserver(v1.NewHTTPHandler(service))
}

func webserver(handler *v1.HTTPHandler) {
	fmt.Println("hei hei")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	if os.Getenv("ENV") == "development" {
		// Human readable logs
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		r.Use(middleware.Logger)
	} else {
		// JSON logs
		logger := httplog.NewLogger("wotapi", httplog.Options{
			JSON: true,
		})
		r.Use(httplog.RequestLogger(logger))
	}

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(server.Addr+"/swagger/doc.json"),
	))
	r.Get("/api/v1/healthz", v1.HealthZ)

	r.Route("/api/v1/manga", func(r chi.Router) {
		r.Get("/", handler.ListManga)
		r.Get("/{mangaID}", handler.GetManga)
		r.Post("/", handler.CreateManga)

		r.Get("/volumes", handler.ListVolumes)
		r.Get("/volume/{volumeID}", handler.GetVolume)
		r.Post("/{mangaID}/volume", handler.CreateVolume)
	})

	err := server.ListenAndServe()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to start server")
	}
}
