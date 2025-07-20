package handler

import "backend/internal/domain"

// Response types for API
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type PaginatedMoviesResponse struct {
	Movies     []*domain.Movie `json:"movies"`
	Page       int             `json:"page"`
	TotalPages int             `json:"totalPages"`
}

type PaginatedTVShowsResponse struct {
	TVShows    []*domain.TVShow `json:"tvShows"`
	Page       int              `json:"page"`
	TotalPages int              `json:"totalPages"`
}

type GenresResponse struct {
	Genres []domain.Genre `json:"genres"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database"`
	Version   string `json:"version"`
}

type WatchlistResponse struct {
	Items      []*domain.Watchlist `json:"items"`
	Total      int64               `json:"total"`
	Page       int                 `json:"page"`
	TotalPages int                 `json:"totalPages"`
}

type AddToWatchlistRequest struct {
	ItemID   string `json:"itemId" binding:"required"`
	ItemType string `json:"itemType" binding:"required,oneof=movie tv"`
}

type RatingRequest struct {
	ItemID   string  `json:"itemId" binding:"required"`
	ItemType string  `json:"itemType" binding:"required,oneof=movie tv"`
	Rating   float64 `json:"rating" binding:"required,min=0,max=10"`
	Review   string  `json:"review"`
}

type RatingsResponse struct {
	Ratings    []*domain.Rating `json:"ratings"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	TotalPages int              `json:"totalPages"`
}

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required,min=3,max=50"`
	FirstName string `json:"firstName" binding:"required,min=1,max=50"`
	LastName  string `json:"lastName" binding:"required,min=1,max=50"`
}

type UpdateUserRequest struct {
	Email     string `json:"email,omitempty" binding:"omitempty,email"`
	Username  string `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
	FirstName string `json:"firstName,omitempty" binding:"omitempty,min=1,max=50"`
	LastName  string `json:"lastName,omitempty" binding:"omitempty,min=1,max=50"`
	Avatar    string `json:"avatar,omitempty"`
}

type UsersResponse struct {
	Users  []*domain.User `json:"users"`
	Total  int64          `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
}
