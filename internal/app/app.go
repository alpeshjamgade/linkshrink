package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"shrinklink/config"
	"shrinklink/internal/clients/cache"
	"shrinklink/internal/clients/db"
	"shrinklink/internal/constants"
	"syscall"

	"github.com/gorilla/mux"

	handler "shrinklink/internal/handler/urls"
	"shrinklink/internal/logger"
	repo "shrinklink/internal/repo/urls"
	service "shrinklink/internal/service/urls"
	"shrinklink/internal/utils"
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
		log.Infof("Starting server on port %s", config.HTTP_PORT)
		http.ListenAndServe(fmt.Sprintf(":%s", config.HTTP_PORT), r)
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
