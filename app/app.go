package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"urlshortner/config"
	"urlshortner/constants"

	handler "urlshortner/handler/urls"
	"urlshortner/logger"
	repo "urlshortner/repo/urls"
	service "urlshortner/service/urls"
	"urlshortner/utils"
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

	urlsRepo := repo.NewUrlsRepo()
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
