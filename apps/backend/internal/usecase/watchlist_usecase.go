package usecase

import (
	"context"
	"time"

	"backend/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchlistUseCase struct {
	watchlistRepo domain.WatchlistRepository
	movieRepo     domain.MovieRepository
	tvShowRepo    domain.TVShowRepository
}

func NewWatchlistUseCase(
	watchlistRepo domain.WatchlistRepository,
	movieRepo domain.MovieRepository,
	tvShowRepo domain.TVShowRepository,
) *WatchlistUseCase {
	return &WatchlistUseCase{
		watchlistRepo: watchlistRepo,
		movieRepo:     movieRepo,
		tvShowRepo:    tvShowRepo,
	}
}

func (uc *WatchlistUseCase) GetUserWatchlist(ctx context.Context, userID string, limit, offset int) ([]*domain.Watchlist, int64, error) {
	return uc.watchlistRepo.GetByUserID(ctx, userID, limit, offset)
}

func (uc *WatchlistUseCase) AddToWatchlist(ctx context.Context, userID, itemID, itemType string) error {
	// Check if item exists
	if itemType == "movie" {
		if _, err := uc.movieRepo.GetByID(ctx, itemID); err != nil {
			return ErrMovieNotFound
		}
	} else if itemType == "tv" {
		if _, err := uc.tvShowRepo.GetByID(ctx, itemID); err != nil {
			return ErrTVShowNotFound
		}
	} else {
		return ErrInvalidInput
	}

	// Check if already in watchlist
	exists, err := uc.watchlistRepo.IsInWatchlist(ctx, userID, itemID, itemType)
	if err != nil {
		return err
	}
	if exists {
		return ErrAlreadyExists
	}

	// Create watchlist item
	watchlist := &domain.Watchlist{
		ID:      primitive.NewObjectID(),
		UserID:  primitive.ObjectID{}, // Convert userID string to ObjectID
		Type:    itemType,
		AddedAt: time.Now(),
	}

	if itemType == "movie" {
		movieObjectID, _ := primitive.ObjectIDFromHex(itemID)
		watchlist.MovieID = &movieObjectID
	} else {
		tvShowObjectID, _ := primitive.ObjectIDFromHex(itemID)
		watchlist.TVShowID = &tvShowObjectID
	}

	return uc.watchlistRepo.Add(ctx, watchlist)
}

func (uc *WatchlistUseCase) RemoveFromWatchlist(ctx context.Context, userID, itemID, itemType string) error {
	// Check if in watchlist
	exists, err := uc.watchlistRepo.IsInWatchlist(ctx, userID, itemID, itemType)
	if err != nil {
		return err
	}
	if !exists {
		return ErrMovieNotFound // Or ErrTVShowNotFound
	}

	return uc.watchlistRepo.Remove(ctx, userID, itemID, itemType)
}

func (uc *WatchlistUseCase) IsInWatchlist(ctx context.Context, userID, itemID, itemType string) (bool, error) {
	return uc.watchlistRepo.IsInWatchlist(ctx, userID, itemID, itemType)
}
