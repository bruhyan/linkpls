package routes

import (
	"net/http"

	"github.com/bruhyan/linkpls/internal/handlers"
)

func Setup(mux *http.ServeMux, lh *handlers.LinkHandler) {
	mux.HandleFunc("/", handlers.HealthHandler)
	mux.HandleFunc("POST /links", lh.CreateShortLinkHandler)
}
