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

	_ "arbolito/docs" // middleware for swagger

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Arbolito API
// @version         1.0
// @description     API for getting different dollar rates in Argentina.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB client
	mongoClient, err := db.NewMongoClient(cfg.MongoURI, cfg.MongoUser, cfg.MongoPassword)
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
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/dolar-blue", rateHandler.GetBlueRate)
	http.HandleFunc("/dolar-oficial", rateHandler.GetOficialRate)
	http.HandleFunc("/dolar-mep", rateHandler.GetMepRate)
	http.HandleFunc("/dolar-ccl", rateHandler.GetCclRate)
	http.HandleFunc("/dolar-tarjeta", rateHandler.GetTarjetaRate)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Start server
	port := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("Server listening on port %s\n", cfg.ServerPort)
	http.ListenAndServe(port, nil)
}
