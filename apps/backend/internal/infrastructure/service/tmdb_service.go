package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"backend/config"
	"backend/internal/domain"
)

type TMDBService struct {
	client  *http.Client
	config  *config.Config
	baseURL string
	apiKey  string
}

func NewTMDBService(cfg *config.Config) *TMDBService {
	return &TMDBService{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		config:  cfg,
		baseURL: cfg.TMDBBaseURL,
		apiKey:  cfg.TMDBAPIKey,
	}
}

// TMDB API response structures
type TMDBMovieResponse struct {
	ID           int         `json:"id"`
	Title        string      `json:"title"`
	Overview     string      `json:"overview"`
	PosterPath   string      `json:"poster_path"`
	BackdropPath string      `json:"backdrop_path"`
	ReleaseDate  string      `json:"release_date"`
	Runtime      int         `json:"runtime"`
	VoteAverage  float64     `json:"vote_average"`
	VoteCount    int         `json:"vote_count"`
	Adult        bool        `json:"adult"`
	Budget       int64       `json:"budget"`
	Revenue      int64       `json:"revenue"`
	Status       string      `json:"status"`
	Tagline      string      `json:"tagline"`
	Genres       []TMDBGenre `json:"genres"`
}

type TMDBTVShowResponse struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Overview         string      `json:"overview"`
	PosterPath       string      `json:"poster_path"`
	BackdropPath     string      `json:"backdrop_path"`
	FirstAirDate     string      `json:"first_air_date"`
	LastAirDate      string      `json:"last_air_date"`
	NumberOfSeasons  int         `json:"number_of_seasons"`
	NumberOfEpisodes int         `json:"number_of_episodes"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int         `json:"vote_count"`
	Status           string      `json:"status"`
	Type             string      `json:"type"`
	Genres           []TMDBGenre `json:"genres"`
}

type TMDBGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TMDBSearchResponse struct {
	Page         int               `json:"page"`
	Results      []json.RawMessage `json:"results"`
	TotalPages   int               `json:"total_pages"`
	TotalResults int               `json:"total_results"`
}

type TMDBGenresResponse struct {
	Genres []TMDBGenre `json:"genres"`
}

func (s *TMDBService) GetMovie(ctx context.Context, movieID int) (*domain.Movie, error) {
	url := fmt.Sprintf("%s/movie/%d?api_key=%s", s.baseURL, movieID, s.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var tmdbMovie TMDBMovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbMovie); err != nil {
		return nil, err
	}

	return s.convertTMDBMovieToMovie(tmdbMovie), nil
}

func (s *TMDBService) GetTVShow(ctx context.Context, tvShowID int) (*domain.TVShow, error) {
	url := fmt.Sprintf("%s/tv/%d?api_key=%s", s.baseURL, tvShowID, s.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var tmdbTVShow TMDBTVShowResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbTVShow); err != nil {
		return nil, err
	}

	return s.convertTMDBTVShowToTVShow(tmdbTVShow), nil
}

func (s *TMDBService) SearchMovies(ctx context.Context, query string, page int) ([]*domain.Movie, int, error) {
	url := fmt.Sprintf("%s/search/movie?api_key=%s&query=%s&page=%d", s.baseURL, s.apiKey, query, page)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var searchResp TMDBSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, 0, err
	}

	var movies []*domain.Movie
	for _, result := range searchResp.Results {
		var tmdbMovie TMDBMovieResponse
		if err := json.Unmarshal(result, &tmdbMovie); err != nil {
			continue // Skip invalid results
		}
		movies = append(movies, s.convertTMDBMovieToMovie(tmdbMovie))
	}

	return movies, searchResp.TotalPages, nil
}

func (s *TMDBService) SearchTVShows(ctx context.Context, query string, page int) ([]*domain.TVShow, int, error) {
	url := fmt.Sprintf("%s/search/tv?api_key=%s&query=%s&page=%d", s.baseURL, s.apiKey, query, page)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var searchResp TMDBSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, 0, err
	}

	var tvShows []*domain.TVShow
	for _, result := range searchResp.Results {
		var tmdbTVShow TMDBTVShowResponse
		if err := json.Unmarshal(result, &tmdbTVShow); err != nil {
			continue // Skip invalid results
		}
		tvShows = append(tvShows, s.convertTMDBTVShowToTVShow(tmdbTVShow))
	}

	return tvShows, searchResp.TotalPages, nil
}

func (s *TMDBService) GetPopularMovies(ctx context.Context, page int) ([]*domain.Movie, int, error) {
	url := fmt.Sprintf("%s/movie/popular?api_key=%s&page=%d", s.baseURL, s.apiKey, page)
	return s.getMoviesList(ctx, url)
}

func (s *TMDBService) GetPopularTVShows(ctx context.Context, page int) ([]*domain.TVShow, int, error) {
	url := fmt.Sprintf("%s/tv/popular?api_key=%s&page=%d", s.baseURL, s.apiKey, page)
	return s.getTVShowsList(ctx, url)
}

func (s *TMDBService) GetMoviesByGenre(ctx context.Context, genreID int, page int) ([]*domain.Movie, int, error) {
	url := fmt.Sprintf("%s/discover/movie?api_key=%s&with_genres=%d&page=%d", s.baseURL, s.apiKey, genreID, page)
	return s.getMoviesList(ctx, url)
}

func (s *TMDBService) GetTVShowsByGenre(ctx context.Context, genreID int, page int) ([]*domain.TVShow, int, error) {
	url := fmt.Sprintf("%s/discover/tv?api_key=%s&with_genres=%d&page=%d", s.baseURL, s.apiKey, genreID, page)
	return s.getTVShowsList(ctx, url)
}

func (s *TMDBService) GetGenres(ctx context.Context, mediaType string) ([]domain.Genre, error) {
	url := fmt.Sprintf("%s/genre/%s/list?api_key=%s", s.baseURL, mediaType, s.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var genresResp TMDBGenresResponse
	if err := json.NewDecoder(resp.Body).Decode(&genresResp); err != nil {
		return nil, err
	}

	var genres []domain.Genre
	for _, g := range genresResp.Genres {
		genres = append(genres, domain.Genre{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	return genres, nil
}

// Helper methods
func (s *TMDBService) getMoviesList(ctx context.Context, url string) ([]*domain.Movie, int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var searchResp TMDBSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, 0, err
	}

	var movies []*domain.Movie
	for _, result := range searchResp.Results {
		var tmdbMovie TMDBMovieResponse
		if err := json.Unmarshal(result, &tmdbMovie); err != nil {
			continue
		}
		movies = append(movies, s.convertTMDBMovieToMovie(tmdbMovie))
	}

	return movies, searchResp.TotalPages, nil
}

func (s *TMDBService) getTVShowsList(ctx context.Context, url string) ([]*domain.TVShow, int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("TMDB API error: %d", resp.StatusCode)
	}

	var searchResp TMDBSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, 0, err
	}

	var tvShows []*domain.TVShow
	for _, result := range searchResp.Results {
		var tmdbTVShow TMDBTVShowResponse
		if err := json.Unmarshal(result, &tmdbTVShow); err != nil {
			continue
		}
		tvShows = append(tvShows, s.convertTMDBTVShowToTVShow(tmdbTVShow))
	}

	return tvShows, searchResp.TotalPages, nil
}

func (s *TMDBService) convertTMDBMovieToMovie(tmdb TMDBMovieResponse) *domain.Movie {
	var genres []domain.Genre
	for _, g := range tmdb.Genres {
		genres = append(genres, domain.Genre{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	return &domain.Movie{
		TMDBMovieID:  tmdb.ID,
		Title:        tmdb.Title,
		Overview:     tmdb.Overview,
		PosterPath:   tmdb.PosterPath,
		BackdropPath: tmdb.BackdropPath,
		ReleaseDate:  tmdb.ReleaseDate,
		Runtime:      tmdb.Runtime,
		VoteAverage:  tmdb.VoteAverage,
		VoteCount:    tmdb.VoteCount,
		Genres:       genres,
		Adult:        tmdb.Adult,
		Budget:       tmdb.Budget,
		Revenue:      tmdb.Revenue,
		Status:       tmdb.Status,
		Tagline:      tmdb.Tagline,
	}
}

func (s *TMDBService) convertTMDBTVShowToTVShow(tmdb TMDBTVShowResponse) *domain.TVShow {
	var genres []domain.Genre
	for _, g := range tmdb.Genres {
		genres = append(genres, domain.Genre{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	return &domain.TVShow{
		TMDBTVShowID:     tmdb.ID,
		Name:             tmdb.Name,
		Overview:         tmdb.Overview,
		PosterPath:       tmdb.PosterPath,
		BackdropPath:     tmdb.BackdropPath,
		FirstAirDate:     tmdb.FirstAirDate,
		LastAirDate:      tmdb.LastAirDate,
		NumberOfSeasons:  tmdb.NumberOfSeasons,
		NumberOfEpisodes: tmdb.NumberOfEpisodes,
		VoteAverage:      tmdb.VoteAverage,
		VoteCount:        tmdb.VoteCount,
		Genres:           genres,
		Status:           tmdb.Status,
		Type:             tmdb.Type,
	}
}
