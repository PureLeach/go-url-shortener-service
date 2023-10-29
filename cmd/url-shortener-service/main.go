package main

import (
	"log/slog"
	"net/http"
	"os"

	"url-shortener-service/internal/config"
	"url-shortener-service/internal/http-server/handlers/redirect"
	"url-shortener-service/internal/http-server/handlers/url/save"
	mwLogger "url-shortener-service/internal/http-server/middleware/logger"
	fLogger "url-shortener-service/internal/lib/logger"
	"url-shortener-service/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.ConfigLoad()

	log := fLogger.SetupLogger(cfg.Env)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "1.0.0"),
	)
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))
	router.Get("/{alias}", redirect.New(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Info("stopping server")

}
