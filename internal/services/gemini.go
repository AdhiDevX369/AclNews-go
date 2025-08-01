package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-test/internal/config"
	"go-test/internal/models"
	"go-test/pkg/errors"
)

// GeminiService handles Gemini AI interactions
type GeminiService struct {
	client *http.Client
	apiKey string
	config *config.Config
}

// GeminiRequest represents the request structure for Gemini API
type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

// GeminiResponse represents the response from Gemini API
type GeminiResponse struct {
	Candidates []GeminiCandidate `json:"candidates"`
}

type GeminiCandidate struct {
	Content GeminiContent `json:"content"`
}

// NewGeminiService creates a new Gemini service
func NewGeminiService(cfg *config.Config) *GeminiService {
	return &GeminiService{
		client: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
		apiKey: cfg.GeminiAPIKey,
		config: cfg,
	}
}

// AnalyzeArticle analyzes a news article using Gemini AI
func (s *GeminiService) AnalyzeArticle(ctx context.Context, article models.Article) (*models.AIAnalysis, error) {
	// Create analysis prompt
	prompt := s.createAnalysisPrompt(article)

	// Prepare request
	request := GeminiRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{Text: prompt},
				},
			},
		},
	}

	// Execute request
	response, err := s.callGeminiAPI(ctx, request)
	if err != nil {
		return nil, err
	}

	// Parse and structure the response
	analysis := s.parseAnalysisResponse(article.ID, response)
	return analysis, nil
}

// BatchAnalyze analyzes multiple articles
func (s *GeminiService) BatchAnalyze(ctx context.Context, articles []models.Article) ([]models.AIAnalysis, error) {
	analyses := make([]models.AIAnalysis, 0, len(articles))

	for _, article := range articles {
		// Rate limiting
		time.Sleep(s.config.RateLimitDelay)

		analysis, err := s.AnalyzeArticle(ctx, article)
		if err != nil {
			// Log error but continue with other articles
			continue
		}

		analyses = append(analyses, *analysis)
	}

	return analyses, nil
}

// createAnalysisPrompt creates a structured prompt for AI analysis
func (s *GeminiService) createAnalysisPrompt(article models.Article) string {
	return fmt.Sprintf(`Analyze this anime/manga news article and provide a structured analysis in the following format:

SUMMARY: [Brief 2-3 sentence summary focusing on the key news]
KEY_POINTS: [List 3-5 key points, separated by semicolons]
SENTIMENT: [POSITIVE/NEUTRAL/NEGATIVE]
RELEVANCE: [Score from 0.0 to 1.0 indicating anime/manga relevance]

Article Details:
Title: %s
Description: %s
Source: %s

Focus on:
- Main anime/manga titles mentioned
- Key developments (new seasons, movies, manga chapters)
- Industry impact
- Fan relevance

Keep the analysis professional and concise.`,
		article.Title,
		article.Description,
		article.Source.Name)
}

// parseAnalysisResponse parses the Gemini response into structured data
func (s *GeminiService) parseAnalysisResponse(articleID, response string) *models.AIAnalysis {
	analysis := &models.AIAnalysis{
		ArticleID:   articleID,
		ProcessedAt: time.Now(),
		Sentiment:   "NEUTRAL",
		Relevance:   0.5,
	}

	// Parse structured response
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "SUMMARY:") {
			analysis.Summary = strings.TrimSpace(strings.TrimPrefix(line, "SUMMARY:"))
		} else if strings.HasPrefix(line, "KEY_POINTS:") {
			keyPointsStr := strings.TrimSpace(strings.TrimPrefix(line, "KEY_POINTS:"))
			analysis.KeyPoints = strings.Split(keyPointsStr, ";")
			// Trim whitespace from each key point
			for i, point := range analysis.KeyPoints {
				analysis.KeyPoints[i] = strings.TrimSpace(point)
			}
		} else if strings.HasPrefix(line, "SENTIMENT:") {
			sentiment := strings.TrimSpace(strings.TrimPrefix(line, "SENTIMENT:"))
			if sentiment == "POSITIVE" || sentiment == "NEGATIVE" || sentiment == "NEUTRAL" {
				analysis.Sentiment = sentiment
			}
		} else if strings.HasPrefix(line, "RELEVANCE:") {
			relevanceStr := strings.TrimSpace(strings.TrimPrefix(line, "RELEVANCE:"))
			if relevance, err := strconv.ParseFloat(relevanceStr, 64); err == nil && relevance >= 0.0 && relevance <= 1.0 {
				analysis.Relevance = relevance
			}
		}
	}

	// Fallback if parsing fails
	if analysis.Summary == "" {
		analysis.Summary = response
	}

	return analysis
}

// callGeminiAPI makes the actual API call to Gemini
func (s *GeminiService) callGeminiAPI(ctx context.Context, request GeminiRequest) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s", s.apiKey)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", errors.Wrap(err, http.StatusInternalServerError, "Failed to marshal Gemini request")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.Wrap(err, http.StatusInternalServerError, "Failed to create Gemini request")
	}

	req.Header.Set("Content-Type", "application/json")

	// Execute with retry logic
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
		return "", errors.Wrap(err, http.StatusServiceUnavailable, "Gemini API request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", errors.ErrInvalidAPIKey
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return "", errors.ErrAPIQuotaExceeded
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.StatusCode, fmt.Sprintf("Gemini API returned status: %d", resp.StatusCode))
	}

	var geminiResp GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return "", errors.Wrap(err, http.StatusInternalServerError, "Failed to decode Gemini response")
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New(http.StatusInternalServerError, "Empty response from Gemini API")
	}

	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}

// ValidateAPIKey validates the Gemini API key
func (s *GeminiService) ValidateAPIKey(ctx context.Context) error {
	testRequest := GeminiRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{Text: "Hello, this is a test."},
				},
			},
		},
	}

	_, err := s.callGeminiAPI(ctx, testRequest)
	return err
}
