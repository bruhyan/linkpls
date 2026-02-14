package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bruhyan/linkpls/internal/service"
)

var linkService *service.LinkService

func CreateShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	fmt.Println("Received URL:", payload.URL)

	shortLink, err := linkService.CreateShortLink(payload.URL)
	if err != nil {
		http.Error(w, "Failed to create short link", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortLink)

}
