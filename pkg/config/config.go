package config

import "os"

type Config struct {
	DolarAPIURL     string
	BluelyticsAPIURL string
	CriptoyaAPIURL  string
	ServerPort      string
}

func LoadConfig() *Config {
	return &Config{
		DolarAPIURL:     getEnv("DOLAR_API_URL", "https://dolarapi.com/v1/dolares/blue"),
		BluelyticsAPIURL: getEnv("BLUELYTICS_API_URL", "https://api.bluelytics.com.ar/v2/latest"),
		CriptoyaAPIURL:  getEnv("CRIPTOYA_API_URL", "https://criptoya.com/api/dolar"),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
