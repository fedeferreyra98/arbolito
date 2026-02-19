package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DolarAPIURL      string
	BluelyticsAPIURL string
	CriptoyaAPIURL   string
	ServerPort       string
	MongoURI         string
	MongoDBName      string
	MongoUser        string
	MongoPassword    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DolarAPIURL:      getEnv("DOLAR_API_URL", "https://dolarapi.com/v1/dolares"),
		BluelyticsAPIURL: getEnv("BLUELYTICS_API_URL", "https://api.bluelytics.com.ar/v2/latest"),
		CriptoyaAPIURL:   getEnv("CRIPTOYA_API_URL", "https://criptoya.com/api/dolar"),
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		MongoURI:         getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName:      getEnv("MONGO_DB_NAME", "arbolito"),
		MongoUser:        getEnv("MONGO_USER", "admin"),
		MongoPassword:    getEnv("MONGO_PASSWORD", "password"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
