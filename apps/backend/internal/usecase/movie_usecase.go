package usecase

import (
	"context"
	"errors"
	"time"

	"backend/internal/domain"
)

var (
	ErrMovieNotFound  = errors.New("movie not found")
	ErrTVShowNotFound = errors.New("tv show not found")
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrAlreadyExists  = errors.New("already exists")
	ErrUnauthorized   = errors.New("unauthorized")
)

type MovieUseCase struct {
	movieRepo   domain.MovieRepository
	tmdbService domain.TMDBService
}

func NewMovieUseCase(movieRepo domain.MovieRepository, tmdbService domain.TMDBService) *MovieUseCase {
	return &MovieUseCase{
		movieRepo:   movieRepo,
		tmdbService: tmdbService,
	}
}

func (uc *MovieUseCase) GetMovieByID(ctx context.Context, id string) (*domain.Movie, error) {
	movie, err := uc.movieRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrMovieNotFound
	}
	return movie, nil
}

func (uc *MovieUseCase) GetMovieByTMDBID(ctx context.Context, tmdbID int) (*domain.Movie, error) {
	// First check if we have it in our database
	movie, err := uc.movieRepo.GetByTMDBID(ctx, tmdbID)
	if err == nil {
		return movie, nil
	}

	// If not found locally, fetch from TMDB and store
	tmdbMovie, err := uc.tmdbService.GetMovie(ctx, tmdbID)
	if err != nil {
		return nil, ErrMovieNotFound
	}

	// Save to our database
	tmdbMovie.CreatedAt = time.Now()
	tmdbMovie.UpdatedAt = time.Now()
	if err := uc.movieRepo.Create(ctx, tmdbMovie); err != nil {
		// Log error but return the movie anyway
		// TODO: Add proper logging
	}

	return tmdbMovie, nil
}

func (uc *MovieUseCase) SearchMovies(ctx context.Context, query string, page int) ([]*domain.Movie, int, error) {
	if query == "" {
		return nil, 0, ErrInvalidInput
	}

	// Get from TMDB
	movies, totalPages, err := uc.tmdbService.SearchMovies(ctx, query, page)
	if err != nil {
		return nil, 0, err
	}

	// Store in our database for future reference
	for _, movie := range movies {
		movie.CreatedAt = time.Now()
		movie.UpdatedAt = time.Now()
		// Check if exists first
		if _, err := uc.movieRepo.GetByTMDBID(ctx, movie.TMDBMovieID); err != nil {
			uc.movieRepo.Create(ctx, movie) // Ignore errors for now
		}
	}

	return movies, totalPages, nil
}

func (uc *MovieUseCase) GetPopularMovies(ctx context.Context, page int) ([]*domain.Movie, int, error) {
	movies, totalPages, err := uc.tmdbService.GetPopularMovies(ctx, page)
	if err != nil {
		return nil, 0, err
	}

	// Store in our database
	for _, movie := range movies {
		movie.CreatedAt = time.Now()
		movie.UpdatedAt = time.Now()
		if _, err := uc.movieRepo.GetByTMDBID(ctx, movie.TMDBMovieID); err != nil {
			uc.movieRepo.Create(ctx, movie)
		}
	}

	return movies, totalPages, nil
}

func (uc *MovieUseCase) GetMoviesByGenre(ctx context.Context, genreID int, page int) ([]*domain.Movie, int, error) {
	movies, totalPages, err := uc.tmdbService.GetMoviesByGenre(ctx, genreID, page)
	if err != nil {
		return nil, 0, err
	}

	// Store in our database
	for _, movie := range movies {
		movie.CreatedAt = time.Now()
		movie.UpdatedAt = time.Now()
		if _, err := uc.movieRepo.GetByTMDBID(ctx, movie.TMDBMovieID); err != nil {
			uc.movieRepo.Create(ctx, movie)
		}
	}

	return movies, totalPages, nil
}
