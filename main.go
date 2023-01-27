package main

import (
	"net/http"
	"os"
	"time"

	v1 "github.com/stianfro/chi-no-wadachi/api/v1"
	_ "github.com/stianfro/chi-no-wadachi/docs"
	utils "github.com/stianfro/chi-no-wadachi/utils"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog/log"
)

// @title Chi No Wadachi API
// @description This is a simple API that is used to demonstrate how to use Chi as a web framework.
// @version 1.0.0
// @license.name MIT
// @BasePath /api/v1
func main() {
	// Environment
	utils.SetEnv()

	// Logger
	logger := httplog.NewLogger("chi-no-wadachi", httplog.Options{
		JSON: true,
	})

	// Service
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(server.Addr+"/swagger/doc.json"),
	))
	r.Get("/api/v1/healthz", v1.HealthZ)

	err := server.ListenAndServe()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to start server")
	}
}
