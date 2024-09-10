package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"linkshrink/clients/cache"
	"linkshrink/clients/db"
	"linkshrink/config"
	"linkshrink/constants"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handler "linkshrink/handler/urls"
	"linkshrink/logger"
	repo "linkshrink/repo/urls"
	service "linkshrink/service/urls"
	"linkshrink/utils"
)

const (
	http_port = "8080"
)

func Start() {

	err := config.LoadConf()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.WithValue(context.Background(), constants.TRACE_ID, utils.GetUUID())
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	log := logger.CreateLoggerWithCtx(ctx)

	db, cache := getClients(ctx)
	urlsRepo := repo.NewUrlsRepo(db, cache)
	urlsService := service.NewUrlsService(urlsRepo)
	urlsHandler := handler.NewUrlsHandler(urlsService)

	r := mux.NewRouter()
	urlsHandler.SetupRoutes(r)
	go func() {
		log.Infof("Starting server on port %s", http_port)
		http.ListenAndServe(fmt.Sprintf(":%s", http_port), r)
	}()

	<-ctx.Done()
	log.Info("Shutting down server...")
}

func getClients(ctx context.Context) (db.DB, cache.ICache) {
	log := logger.CreateLoggerWithCtx(ctx)

	db := db.NewPostgresDB(config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME)
	err := db.Connect(ctx)
	if err != nil {
		log.Errorf("Error connecting to database: %v", err)
	}

	cache := cache.NewRedisCache(config.REDIS_HOST, config.REDIS_PORT)
	if err := cache.Connect(ctx); err != nil {
		log.Errorf("Error connecting to database: %v", err)
	}

	return db, cache
}
