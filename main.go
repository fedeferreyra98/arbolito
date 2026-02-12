package main

import (
	"arbolito/pkg/adapter/bluelytics"
	"arbolito/pkg/adapter/criptoya"
	"arbolito/pkg/adapter/dolarapi"
	"arbolito/pkg/config"
	"arbolito/pkg/handler"
	"arbolito/pkg/repository"
	"arbolito/pkg/service"
	"fmt"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize adapters
	dolarAPI := dolarapi.NewDolarAPI(cfg.DolarAPIURL)
	bluelyticsAPI := bluelytics.NewBluelyticsAPI(cfg.BluelyticsAPIURL)
	criptoyaAPI := criptoya.NewCriptoyaAPI(cfg.CriptoyaAPIURL)

	// Create a list of repositories
	repos := []repository.RateRepository{
		dolarAPI,
		bluelyticsAPI,
		criptoyaAPI,
	}

	// Initialize service
	rateService := service.NewRateService(repos)

	// Initialize handler
	rateHandler := handler.NewRateHandler(rateService)

	// Define routes
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/dolar-blue", rateHandler.GetAverageRate)

	// Start server
	port := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("Server listening on port %s\n", cfg.ServerPort)
	http.ListenAndServe(port, nil)
}
