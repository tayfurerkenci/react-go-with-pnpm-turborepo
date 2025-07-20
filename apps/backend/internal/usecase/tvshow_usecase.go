package usecase

import (
	"context"
	"time"

	"backend/internal/domain"
)

type TVShowUseCase struct {
	tvShowRepo  domain.TVShowRepository
	tmdbService domain.TMDBService
}

func NewTVShowUseCase(tvShowRepo domain.TVShowRepository, tmdbService domain.TMDBService) *TVShowUseCase {
	return &TVShowUseCase{
		tvShowRepo:  tvShowRepo,
		tmdbService: tmdbService,
	}
}

func (uc *TVShowUseCase) GetTVShowByID(ctx context.Context, id string) (*domain.TVShow, error) {
	tvShow, err := uc.tvShowRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrTVShowNotFound
	}
	return tvShow, nil
}

func (uc *TVShowUseCase) GetTVShowByTMDBID(ctx context.Context, tmdbID int) (*domain.TVShow, error) {
	// First check if we have it in our database
	tvShow, err := uc.tvShowRepo.GetByTMDBID(ctx, tmdbID)
	if err == nil {
		return tvShow, nil
	}

	// If not found locally, fetch from TMDB and store
	tmdbTVShow, err := uc.tmdbService.GetTVShow(ctx, tmdbID)
	if err != nil {
		return nil, ErrTVShowNotFound
	}

	// Save to our database
	tmdbTVShow.CreatedAt = time.Now()
	tmdbTVShow.UpdatedAt = time.Now()
	if err := uc.tvShowRepo.Create(ctx, tmdbTVShow); err != nil {
		// Log error but return the TV show anyway
	}

	return tmdbTVShow, nil
}

func (uc *TVShowUseCase) SearchTVShows(ctx context.Context, query string, page int) ([]*domain.TVShow, int, error) {
	if query == "" {
		return nil, 0, ErrInvalidInput
	}

	// Get from TMDB
	tvShows, totalPages, err := uc.tmdbService.SearchTVShows(ctx, query, page)
	if err != nil {
		return nil, 0, err
	}

	// Store in our database for future reference
	for _, tvShow := range tvShows {
		tvShow.CreatedAt = time.Now()
		tvShow.UpdatedAt = time.Now()
		// Check if exists first
		if _, err := uc.tvShowRepo.GetByTMDBID(ctx, tvShow.TMDBTVShowID); err != nil {
			uc.tvShowRepo.Create(ctx, tvShow) // Ignore errors for now
		}
	}

	return tvShows, totalPages, nil
}

func (uc *TVShowUseCase) GetPopularTVShows(ctx context.Context, page int) ([]*domain.TVShow, int, error) {
	tvShows, totalPages, err := uc.tmdbService.GetPopularTVShows(ctx, page)
	if err != nil {
		return nil, 0, err
	}

	// Store in our database
	for _, tvShow := range tvShows {
		tvShow.CreatedAt = time.Now()
		tvShow.UpdatedAt = time.Now()
		if _, err := uc.tvShowRepo.GetByTMDBID(ctx, tvShow.TMDBTVShowID); err != nil {
			uc.tvShowRepo.Create(ctx, tvShow)
		}
	}

	return tvShows, totalPages, nil
}

func (uc *TVShowUseCase) GetTVShowsByGenre(ctx context.Context, genreID int, page int) ([]*domain.TVShow, int, error) {
	tvShows, totalPages, err := uc.tmdbService.GetTVShowsByGenre(ctx, genreID, page)
	if err != nil {
		return nil, 0, err
	}

	// Store in our database
	for _, tvShow := range tvShows {
		tvShow.CreatedAt = time.Now()
		tvShow.UpdatedAt = time.Now()
		if _, err := uc.tvShowRepo.GetByTMDBID(ctx, tvShow.TMDBTVShowID); err != nil {
			uc.tvShowRepo.Create(ctx, tvShow)
		}
	}

	return tvShows, totalPages, nil
}
