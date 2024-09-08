package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	handler "urlshortner/handler/urls"
	repo "urlshortner/repo/urls"
	service "urlshortner/service/urls"
)

const (
	http_port = "8080"
)

func Start() {
	urlsRepo := repo.NewUrlsRepo()
	urlsService := service.NewUrlsService(urlsRepo)
	urlsHandler := handler.NewUrlsHandler(urlsService)

	r := mux.NewRouter()
	urlsHandler.SetupRoutes(r)
	err := http.ListenAndServe(fmt.Sprintf(":%s", http_port), r)
	if err != nil {
		println("Error starting server", "error", err)
		return
	}
	println("Starting server", "port", http_port)

}
