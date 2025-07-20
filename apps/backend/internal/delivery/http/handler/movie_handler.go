package handler

import (
	"net/http"
	"strconv"

	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUseCase *usecase.MovieUseCase
}

func NewMovieHandler(movieUseCase *usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{
		movieUseCase: movieUseCase,
	}
}

// GetMovie godoc
// @Summary Get movie by ID
// @Description Get movie details by internal ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} domain.Movie
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovie(c *gin.Context) {
	id := c.Param("id")

	movie, err := h.movieUseCase.GetMovieByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "movie_not_found",
			Message: "Movie not found",
		})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// GetMovieByTMDBID godoc
// @Summary Get movie by TMDB ID
// @Description Get movie details by TMDB ID (fetches from TMDB if not in database)
// @Tags movies
// @Accept json
// @Produce json
// @Param tmdb_id path int true "TMDB Movie ID"
// @Success 200 {object} domain.Movie
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /movies/tmdb/{tmdb_id} [get]
func (h *MovieHandler) GetMovieByTMDBID(c *gin.Context) {
	tmdbIDStr := c.Param("tmdb_id")
	tmdbID, err := strconv.Atoi(tmdbIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_tmdb_id",
			Message: "Invalid TMDB ID format",
		})
		return
	}

	movie, err := h.movieUseCase.GetMovieByTMDBID(c.Request.Context(), tmdbID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "movie_not_found",
			Message: "Movie not found",
		})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// SearchMovies godoc
// @Summary Search movies
// @Description Search movies by title or overview
// @Tags movies
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Success 200 {object} PaginatedMoviesResponse
// @Failure 400 {object} ErrorResponse
// @Router /movies/search [get]
func (h *MovieHandler) SearchMovies(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "missing_query",
			Message: "Search query is required",
		})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	movies, totalPages, err := h.movieUseCase.SearchMovies(c.Request.Context(), query, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "search_error",
			Message: "Failed to search movies",
		})
		return
	}

	c.JSON(http.StatusOK, PaginatedMoviesResponse{
		Movies:     movies,
		Page:       page,
		TotalPages: totalPages,
	})
}

// GetPopularMovies godoc
// @Summary Get popular movies
// @Description Get list of popular movies from TMDB
// @Tags movies
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Success 200 {object} PaginatedMoviesResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies/popular [get]
func (h *MovieHandler) GetPopularMovies(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	movies, totalPages, err := h.movieUseCase.GetPopularMovies(c.Request.Context(), page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "fetch_error",
			Message: "Failed to fetch popular movies",
		})
		return
	}

	c.JSON(http.StatusOK, PaginatedMoviesResponse{
		Movies:     movies,
		Page:       page,
		TotalPages: totalPages,
	})
}

// GetMoviesByGenre godoc
// @Summary Get movies by genre
// @Description Get movies filtered by genre
// @Tags movies
// @Accept json
// @Produce json
// @Param genre_id path int true "Genre ID"
// @Param page query int false "Page number" default(1)
// @Success 200 {object} PaginatedMoviesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies/genre/{genre_id} [get]
func (h *MovieHandler) GetMoviesByGenre(c *gin.Context) {
	genreIDStr := c.Param("genre_id")
	genreID, err := strconv.Atoi(genreIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_genre_id",
			Message: "Invalid genre ID format",
		})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	movies, totalPages, err := h.movieUseCase.GetMoviesByGenre(c.Request.Context(), genreID, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "fetch_error",
			Message: "Failed to fetch movies by genre",
		})
		return
	}

	c.JSON(http.StatusOK, PaginatedMoviesResponse{
		Movies:     movies,
		Page:       page,
		TotalPages: totalPages,
	})
}
