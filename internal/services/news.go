package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"go-test/internal/config"
	"go-test/internal/models"
	"go-test/pkg/errors"
)

// NewsService handles news API interactions
type NewsService struct {
	client *http.Client
	apiKey string
	config *config.Config
}

// NewNewsService creates a new news service
func NewNewsService(cfg *config.Config) *NewsService {
	return &NewsService{
		client: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
		apiKey: cfg.NewsAPIKey,
		config: cfg,
	}
}

// FetchAnimeNews fetches anime-related news articles
func (s *NewsService) FetchAnimeNews(ctx context.Context) ([]models.Article, error) {
	// Build URL with parameters
	baseURL := "https://newsapi.org/v2/everything"
	params := url.Values{
		"q":        {"anime OR manga OR otaku"},
		"sortBy":   {"publishedAt"},
		"pageSize": {fmt.Sprintf("%d", s.config.MaxArticles)},
		"language": {"en"},
		"apiKey":   {s.apiKey},
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, http.StatusInternalServerError, "Failed to create news request")
	}

	// Set headers
	req.Header.Set("User-Agent", "AnimeNewsAI/1.0")
	req.Header.Set("Accept", "application/json")

	// Execute request with retry logic
	var resp *http.Response
	for attempt := 0; attempt < s.config.RetryAttempts; attempt++ {
		resp, err = s.client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			break
		}

		if attempt < s.config.RetryAttempts-1 {
			time.Sleep(s.config.RateLimitDelay * time.Duration(attempt+1))
		}
	}

	if err != nil {
		return nil, errors.Wrap(err, http.StatusServiceUnavailable, "News API request failed")
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.ErrInvalidAPIKey
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, errors.ErrAPIQuotaExceeded
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.StatusCode, fmt.Sprintf("News API returned status: %d", resp.StatusCode))
	}

	// Parse response
	var newsResp struct {
		Status       string `json:"status"`
		TotalResults int    `json:"totalResults"`
		Articles     []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Content     string `json:"content"`
			URL         string `json:"url"`
			URLToImage  string `json:"urlToImage"`
			PublishedAt string `json:"publishedAt"`
			Source      struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"source"`
			Author string `json:"author"`
		} `json:"articles"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&newsResp); err != nil {
		return nil, errors.Wrap(err, http.StatusInternalServerError, "Failed to decode news response")
	}

	// Convert to internal models
	articles := make([]models.Article, 0, len(newsResp.Articles))
	for i, article := range newsResp.Articles {
		publishedAt, _ := time.Parse(time.RFC3339, article.PublishedAt)

		articles = append(articles, models.Article{
			ID:          fmt.Sprintf("news_%d_%d", time.Now().Unix(), i),
			Title:       article.Title,
			Description: article.Description,
			Content:     article.Content,
			URL:         article.URL,
			URLToImage:  article.URLToImage,
			PublishedAt: publishedAt,
			Source: models.Source{
				ID:   article.Source.ID,
				Name: article.Source.Name,
			},
			Author: article.Author,
		})
	}

	return articles, nil
}

// ValidateAPIKey validates the news API key
func (s *NewsService) ValidateAPIKey(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&pageSize=1&apiKey=%s", s.apiKey), nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.ErrInvalidAPIKey
	}

	return nil
}
