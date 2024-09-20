package urls

import (
	"context"
	"net/http"
	"shrinklink/constants"
	"shrinklink/logger"
	"shrinklink/utils"

	"github.com/gorilla/mux"
)

func (h *UrlsHandler) GetAllUrls(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TRACE_ID, utils.GetUUID())

	result, err := h.service.GetAllUrls(ctx)

	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := utils.HTTPResponse{Data: result, Status: "success", Message: ""}

	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *UrlsHandler) AddUrl(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TRACE_ID, utils.GetUUID())
	log := logger.CreateLoggerWithCtx(ctx)
	type requestPayload struct {
		Url string `json:"url"`
	}
	req := &requestPayload{}
	res := utils.HTTPResponse{Data: map[string]string{}, Status: "success", Message: ""}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		log.Errorw("Error while reading request", "error", err.Error())
		res.Status = "error"
		res.Message = "Invalid request"
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	shortUrl, err := h.service.AddUrl(ctx, req.Url)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	response := utils.HTTPResponse{}
	response.Status = http.StatusText(http.StatusOK)
	response.Data = map[string]string{"short_url": shortUrl}
	response.Message = "Successfully generated short url"
	utils.WriteJSON(w, http.StatusOK, response)
	return

}

func (h *UrlsHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TRACE_ID, utils.GetUUID())
	res := utils.HTTPResponse{Data: map[string]string{}, Status: "success", Message: ""}
	req := mux.Vars(r)
	url, err := h.service.GetUrlWithHash(ctx, req["short_url"])
	if err != nil {
		res.Status = "error"
		res.Message = "Error while fetching url"
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}
	//utils.WriteJSON(w, http.StatusTemporaryRedirect, res, map[string][]string{"Location": {url}})
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return

}
