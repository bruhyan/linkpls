package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bruhyan/linkpls/internal/database"
	"github.com/bruhyan/linkpls/internal/handlers"
	"github.com/bruhyan/linkpls/internal/routes"
	"github.com/bruhyan/linkpls/internal/service"
)

func main() {
	PORT := "8080"
	mux := http.NewServeMux()

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisClient := database.New(redisAddr, "", 0, 0)

	if err := redisClient.Ping(context.Background()); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	linkService := service.New(redisClient)
	linkHandler := handlers.NewLinkHandler(linkService)
	routes.Setup(mux, linkHandler)

	fmt.Println("linkpls server running on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
