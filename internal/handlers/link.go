package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bruhyan/linkpls/internal/service"
)

type LinkHandler struct {
	service *service.LinkService
}

func NewLinkHandler(service *service.LinkService) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

func (h *LinkHandler) CreateShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	fmt.Println("Received URL:", payload.URL)

	shortLink, err := h.service.CreateShortLink(payload.URL)
	if err != nil {
		http.Error(w, "Failed to create short link", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortLink)

}
