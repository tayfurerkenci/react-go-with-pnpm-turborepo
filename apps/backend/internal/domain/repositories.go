package domain

import "context"

// MovieRepository defines movie data access interface
type MovieRepository interface {
	GetByID(ctx context.Context, id string) (*Movie, error)
	GetByTMDBID(ctx context.Context, tmdbID int) (*Movie, error)
	Create(ctx context.Context, movie *Movie) error
	Update(ctx context.Context, movie *Movie) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*Movie, int64, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*Movie, int64, error)
	GetByGenre(ctx context.Context, genreID int, limit, offset int) ([]*Movie, int64, error)
	GetPopular(ctx context.Context, limit, offset int) ([]*Movie, int64, error)
}

// TVShowRepository defines TV show data access interface
type TVShowRepository interface {
	GetByID(ctx context.Context, id string) (*TVShow, error)
	GetByTMDBID(ctx context.Context, tmdbID int) (*TVShow, error)
	Create(ctx context.Context, tvShow *TVShow) error
	Update(ctx context.Context, tvShow *TVShow) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*TVShow, int64, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*TVShow, int64, error)
	GetByGenre(ctx context.Context, genreID int, limit, offset int) ([]*TVShow, int64, error)
	GetPopular(ctx context.Context, limit, offset int) ([]*TVShow, int64, error)
}

// UserRepository defines user data access interface
type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*User, int64, error)
}

// WatchlistRepository defines watchlist data access interface
type WatchlistRepository interface {
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*Watchlist, int64, error)
	Add(ctx context.Context, watchlist *Watchlist) error
	Remove(ctx context.Context, userID, itemID string, itemType string) error
	IsInWatchlist(ctx context.Context, userID, itemID string, itemType string) (bool, error)
}

// RatingRepository defines rating data access interface
type RatingRepository interface {
	GetByUserAndItem(ctx context.Context, userID, itemID string, itemType string) (*Rating, error)
	GetByItem(ctx context.Context, itemID string, itemType string, limit, offset int) ([]*Rating, int64, error)
	GetByUser(ctx context.Context, userID string, limit, offset int) ([]*Rating, int64, error)
	Create(ctx context.Context, rating *Rating) error
	Update(ctx context.Context, rating *Rating) error
	Delete(ctx context.Context, id string) error
	GetAverageRating(ctx context.Context, itemID string, itemType string) (float64, int64, error)
}

// TMDBService defines external TMDB API interface
type TMDBService interface {
	GetMovie(ctx context.Context, movieID int) (*Movie, error)
	GetTVShow(ctx context.Context, tvShowID int) (*TVShow, error)
	SearchMovies(ctx context.Context, query string, page int) ([]*Movie, int, error)
	SearchTVShows(ctx context.Context, query string, page int) ([]*TVShow, int, error)
	GetPopularMovies(ctx context.Context, page int) ([]*Movie, int, error)
	GetPopularTVShows(ctx context.Context, page int) ([]*TVShow, int, error)
	GetMoviesByGenre(ctx context.Context, genreID int, page int) ([]*Movie, int, error)
	GetTVShowsByGenre(ctx context.Context, genreID int, page int) ([]*TVShow, int, error)
	GetGenres(ctx context.Context, mediaType string) ([]Genre, error)
}
