package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie represents a movie entity
type Movie struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TMDBMovieID  int                `json:"tmdbMovieId" bson:"tmdbMovieId"`
	Title        string             `json:"title" bson:"title"`
	Overview     string             `json:"overview" bson:"overview"`
	PosterPath   string             `json:"posterPath" bson:"posterPath"`
	BackdropPath string             `json:"backdropPath" bson:"backdropPath"`
	ReleaseDate  string             `json:"releaseDate" bson:"releaseDate"`
	Runtime      int                `json:"runtime" bson:"runtime"`
	VoteAverage  float64            `json:"voteAverage" bson:"voteAverage"`
	VoteCount    int                `json:"voteCount" bson:"voteCount"`
	Genres       []Genre            `json:"genres" bson:"genres"`
	Adult        bool               `json:"adult" bson:"adult"`
	Budget       int64              `json:"budget" bson:"budget"`
	Revenue      int64              `json:"revenue" bson:"revenue"`
	Status       string             `json:"status" bson:"status"`
	Tagline      string             `json:"tagline" bson:"tagline"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// TVShow represents a TV show entity
type TVShow struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TMDBTVShowID     int                `json:"tmdbTvShowId" bson:"tmdbTvShowId"`
	Name             string             `json:"name" bson:"name"`
	Overview         string             `json:"overview" bson:"overview"`
	PosterPath       string             `json:"posterPath" bson:"posterPath"`
	BackdropPath     string             `json:"backdropPath" bson:"backdropPath"`
	FirstAirDate     string             `json:"firstAirDate" bson:"firstAirDate"`
	LastAirDate      string             `json:"lastAirDate" bson:"lastAirDate"`
	NumberOfSeasons  int                `json:"numberOfSeasons" bson:"numberOfSeasons"`
	NumberOfEpisodes int                `json:"numberOfEpisodes" bson:"numberOfEpisodes"`
	VoteAverage      float64            `json:"voteAverage" bson:"voteAverage"`
	VoteCount        int                `json:"voteCount" bson:"voteCount"`
	Genres           []Genre            `json:"genres" bson:"genres"`
	Status           string             `json:"status" bson:"status"`
	Type             string             `json:"type" bson:"type"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Genre represents a genre entity
type Genre struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// User represents a user entity
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Username  string             `json:"username" bson:"username"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Avatar    string             `json:"avatar" bson:"avatar"`
	IsActive  bool               `json:"isActive" bson:"isActive"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Watchlist represents a user's watchlist
type Watchlist struct {
	ID       primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID  `json:"userId" bson:"userId"`
	MovieID  *primitive.ObjectID `json:"movieId,omitempty" bson:"movieId,omitempty"`
	TVShowID *primitive.ObjectID `json:"tvShowId,omitempty" bson:"tvShowId,omitempty"`
	Type     string              `json:"type" bson:"type"` // "movie" or "tv"
	AddedAt  time.Time           `json:"addedAt" bson:"addedAt"`
}

// Rating represents a user's rating for a movie or TV show
type Rating struct {
	ID        primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID  `json:"userId" bson:"userId"`
	MovieID   *primitive.ObjectID `json:"movieId,omitempty" bson:"movieId,omitempty"`
	TVShowID  *primitive.ObjectID `json:"tvShowId,omitempty" bson:"tvShowId,omitempty"`
	Type      string              `json:"type" bson:"type"` // "movie" or "tv"
	Rating    float64             `json:"rating" bson:"rating"`
	Review    string              `json:"review" bson:"review"`
	CreatedAt time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt" bson:"updatedAt"`
}
