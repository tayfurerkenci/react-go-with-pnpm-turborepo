package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	MongoURI    string
	Environment string
	JWTSecret   string
	TMDBAPIKey  string
	TMDBBaseURL string
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017/movies_platform"),
		Environment: getEnv("ENVIRONMENT", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		TMDBAPIKey:  getEnv("TMDB_API_KEY", "e2cadf69ee384867df4db7959f1eee53"),
		TMDBBaseURL: getEnv("TMDB_BASE_URL", "https://api.themoviedb.org/3"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
