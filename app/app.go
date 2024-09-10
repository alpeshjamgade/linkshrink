package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"shrink-link/clients/cache"
	"shrink-link/clients/db"
	"shrink-link/config"
	"shrink-link/constants"
	"syscall"

	"github.com/gorilla/mux"

	handler "shrink-link/handler/urls"
	"shrink-link/logger"
	repo "shrink-link/repo/urls"
	service "shrink-link/service/urls"
	"shrink-link/utils"
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
