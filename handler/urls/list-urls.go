package urls

import (
	"context"
	"net/http"
	"urlshortner/constants"
	"urlshortner/utils"
)

func (h *UrlsHandler) ListUrls(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TRACE_ID, utils.GetUUID())

	data, err := h.service.ListUrls(ctx)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
	}

	payload := utils.HTTPResponse{
		Data:   data,
		Status: http.StatusText(http.StatusOK),
	}

	utils.WriteJSON(w, http.StatusOK, payload)

}
