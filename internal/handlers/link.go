package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

}
