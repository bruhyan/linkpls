package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bruhyan/linkpls/internal/routes"
)

func main() {
	PORT := "8080"
	mux := http.NewServeMux()
	routes.Setup(mux)

	fmt.Println("linkpls server running on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
