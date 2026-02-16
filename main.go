package main

import (
	"arbolito/pkg/adapter/bluelytics"
	"arbolito/pkg/adapter/criptoya"
	"arbolito/pkg/adapter/dolarapi"
	"arbolito/pkg/config"
	"arbolito/pkg/db"
	"arbolito/pkg/handler"
	"arbolito/pkg/repository"
	"arbolito/pkg/repository/caching"
	"arbolito/pkg/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB client
	mongoClient, err := db.NewMongoClient(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	mongoDB := mongoClient.Database(cfg.MongoDBName)

	// Initialize caching repository
	cachingRepo, err := caching.NewMongoCachingRepository(mongoDB)
	if err != nil {
		log.Fatalf("Failed to create caching repository: %v", err)
	}

	// Initialize adapters
	dolarAPI := dolarapi.NewDolarAPIAdapter(cfg.DolarAPIURL)
	bluelyticsAPI := bluelytics.NewBluelyticsAdapter(cfg.BluelyticsAPIURL)
	criptoyaAPI := criptoya.NewCriptoyaAdapter(cfg.CriptoyaAPIURL)

	// Create a list of repositories
	repos := []repository.RateApiAdapter{
		dolarAPI,
		bluelyticsAPI,
		criptoyaAPI,
	}

	// Initialize service
	rateService := service.NewRateService(repos, cachingRepo)

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
