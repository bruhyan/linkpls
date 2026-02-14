package routes

import (
	"net/http"

	"github.com/bruhyan/linkpls/internal/handlers"
)

func Setup(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HealthHandler)
	mux.HandleFunc("POST /links", handlers.CreateShortLinkHandler)
}
