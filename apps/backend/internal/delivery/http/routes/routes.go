package routes

import (
	"backend/config"
	"backend/internal/delivery/http/handler"
	"backend/internal/infrastructure/repository"
	"backend/internal/infrastructure/service"
	"backend/internal/middleware"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.RouterGroup, db *mongo.Database, cfg *config.Config) {
	// Initialize services
	tmdbService := service.NewTMDBService(cfg)

	// Initialize repositories
	movieRepo := repository.NewMovieRepository(db)
	// tvShowRepo := repository.NewTVShowRepository(db)
	// userRepo := repository.NewUserRepository(db)
	// watchlistRepo := repository.NewWatchlistRepository(db)

	// Initialize use cases
	movieUseCase := usecase.NewMovieUseCase(movieRepo, tmdbService)
	// tvShowUseCase := usecase.NewTVShowUseCase(tvShowRepo, tmdbService)
	// watchlistUseCase := usecase.NewWatchlistUseCase(watchlistRepo, movieRepo, tvShowRepo)

	// Initialize handlers
	movieHandler := handler.NewMovieHandler(movieUseCase)
	// tvShowHandler := handler.NewTVShowHandler(tvShowUseCase)
	// watchlistHandler := handler.NewWatchlistHandler(watchlistUseCase)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, handler.HealthResponse{
			Status:    "healthy",
			Timestamp: "2025-01-20T23:00:00Z",
			Database:  "connected",
			Version:   "1.0.0",
		})
	})

	// API v1 routes
	v1 := router.Group("/v1")
	v1.Use(middleware.RequestID())
	v1.Use(middleware.ErrorHandler())

	// Movie routes
	movies := v1.Group("/movies")
	{
		movies.GET("/:id", movieHandler.GetMovie)
		movies.GET("/tmdb/:tmdb_id", movieHandler.GetMovieByTMDBID)
		movies.GET("/search", movieHandler.SearchMovies)
		movies.GET("/popular", movieHandler.GetPopularMovies)
		movies.GET("/genre/:genre_id", movieHandler.GetMoviesByGenre)
	}

	// TV Show routes (placeholder)
	tvShows := v1.Group("/tv")
	{
		tvShows.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "TV show endpoint - coming soon"})
		})
	}

	// Watchlist routes (placeholder)
	watchlist := v1.Group("/watchlist")
	{
		watchlist.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Watchlist endpoint - coming soon"})
		})
	}

	// Genres endpoint
	v1.GET("/genres/:media_type", func(c *gin.Context) {
		mediaType := c.Param("media_type")
		if mediaType != "movie" && mediaType != "tv" {
			c.JSON(400, handler.ErrorResponse{
				Error:   "invalid_media_type",
				Message: "Media type must be 'movie' or 'tv'",
			})
			return
		}

		genres, err := tmdbService.GetGenres(c.Request.Context(), mediaType)
		if err != nil {
			c.JSON(500, handler.ErrorResponse{
				Error:   "fetch_error",
				Message: "Failed to fetch genres",
			})
			return
		}

		c.JSON(200, handler.GenresResponse{
			Genres: genres,
		})
	})
}
