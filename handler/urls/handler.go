package urls

import (
	"net/http"
	"shrinklink/middlewares"
	"shrinklink/service/urls"

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
	r.HandleFunc("/urls", h.GetAllUrls).Methods(http.MethodGet)
	r.HandleFunc("/urls", h.AddUrl).Methods(http.MethodPost)
	r.HandleFunc("/{short_url}", h.GetUrl).Methods(http.MethodGet)
}
