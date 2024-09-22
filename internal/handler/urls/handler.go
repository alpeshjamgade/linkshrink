package urls

import (
	"net/http"
	"shrinklink/internal/middlewares"
	"shrinklink/internal/service/urls"

	"github.com/gorilla/mux"
)

type UrlsHandler struct {
	service urls.IUrlService
}

func NewUrlsHandler(service urls.IUrlService) *UrlsHandler {
	handler := &UrlsHandler{service: service}
	return handler
}

func (h *UrlsHandler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.EnableCORS)
	r.HandleFunc("/api/urls", h.GetAllUrls).Methods(http.MethodGet)
	r.HandleFunc("/api/urls", h.AddUrl).Methods(http.MethodPost)
	r.HandleFunc("/api/{short_url}", h.GetUrl).Methods(http.MethodGet)
}
