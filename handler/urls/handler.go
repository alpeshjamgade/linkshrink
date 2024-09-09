package urls

import (
	"github.com/gorilla/mux"
	"net/http"
	"urlshortner/service/urls"
)

type UrlsHandler struct {
	service urls.IUrlService
}

func NewUrlsHandler(service urls.IUrlService) *UrlsHandler {
	handler := &UrlsHandler{service: service}
	return handler
}

func (h *UrlsHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/urls", h.ListUrls).Methods(http.MethodGet)
	r.HandleFunc("/urls", h.AddUrl).Methods(http.MethodPost)
	r.HandleFunc("/url/{short_url}", h.GetUrl).Methods(http.MethodGet)
}
