package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-test/internal/config"
	"go-test/internal/models"
	"go-test/internal/services"
	"go-test/pkg/logger"
)

const (
	AppName    = "Anime News AI"
	AppVersion = "1.0.0"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	appLogger := logger.New(cfg.LogLevel, cfg.LogFormat)
	defer appLogger.Close()

	// Log startup
	appLogger.WithFields(map[string]interface{}{
		"app":     AppName,
		"version": AppVersion,
		"env":     cfg.Environment,
	}).Info("Starting application")

	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize services
	newsService := services.NewNewsService(cfg)
	geminiService := services.NewGeminiService(cfg)

	// Validate API keys
	appLogger.Info("Validating API keys...")
	if err := validateServices(ctx, newsService, geminiService); err != nil {
		appLogger.WithError(err).Fatal("Service validation failed")
	}
	appLogger.Info("API keys validated successfully")

	// Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Run main application logic
	go func() {
		if err := runApplication(ctx, cfg, newsService, geminiService, appLogger); err != nil {
			appLogger.WithError(err).Error("Application error")
			cancel()
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	appLogger.Info("Shutdown signal received, gracefully shutting down...")

	// Cancel context to stop all operations
	cancel()

	// Give operations time to clean up
	time.Sleep(2 * time.Second)
	appLogger.Info("Application stopped")
}

// validateServices validates that all external services are accessible
func validateServices(ctx context.Context, newsService *services.NewsService, geminiService *services.GeminiService) error {
	// Validate News API
	if err := newsService.ValidateAPIKey(ctx); err != nil {
		return fmt.Errorf("news API validation failed: %w", err)
	}

	// Validate Gemini API
	if err := geminiService.ValidateAPIKey(ctx); err != nil {
		return fmt.Errorf("gemini API validation failed: %w", err)
	}

	return nil
}

// runApplication contains the main application logic
func runApplication(ctx context.Context, cfg *config.Config, newsService *services.NewsService, geminiService *services.GeminiService, logger *logger.Logger) error {
	logger.Info("🍃 Anime News AI Assistant Started")
	logger.Info("==================================")

	// Fetch anime news
	logger.Info("📰 Fetching latest anime news...")
	articles, err := newsService.FetchAnimeNews(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch news: %w", err)
	}

	if len(articles) == 0 {
		logger.Warn("No anime news articles found")
		return nil
	}

	logger.WithField("count", len(articles)).Info("Articles fetched successfully")

	// Process articles with AI
	logger.Info("🤖 Processing articles with Gemini AI...")

	for i, article := range articles {
		select {
		case <-ctx.Done():
			logger.Info("Processing cancelled")
			return ctx.Err()
		default:
		}

		logger.WithFields(map[string]interface{}{
			"article": i + 1,
			"total":   len(articles),
			"title":   article.Title,
		}).Info("Processing article")

		// Analyze with Gemini
		analysis, err := geminiService.AnalyzeArticle(ctx, article)
		if err != nil {
			logger.WithFields(map[string]interface{}{
				"article_id": article.ID,
				"title":      article.Title,
			}).WithError(err).Error("Failed to analyze article")
			continue
		}

		// Display results
		displayArticleAnalysis(article, analysis, logger)

		// Rate limiting
		time.Sleep(cfg.RateLimitDelay)
	}

	logger.Info("✅ Processing complete!")
	return nil
}

// displayArticleAnalysis displays the article and its AI analysis
func displayArticleAnalysis(article models.Article, analysis *models.AIAnalysis, logger *logger.Logger) {
	// Log structured data for monitoring
	fields := map[string]interface{}{
		"article_title":     article.Title,
		"article_source":    article.Source.Name,
		"article_published": article.PublishedAt.Format("2006-01-02 15:04"),
		"article_url":       article.URL,
	}

	if analysis != nil {
		fields["ai_summary"] = analysis.Summary
		fields["ai_sentiment"] = analysis.Sentiment
		fields["ai_key_points_count"] = len(analysis.KeyPoints)

		logger.WithFields(fields).Info("Article processed successfully")

		// Display user-friendly output
		logger.Info("📄 Article: " + article.Title)
		logger.Info("🧠 AI Summary: " + analysis.Summary)
		logger.Info("📊 Sentiment: " + analysis.Sentiment)

		if len(analysis.KeyPoints) > 0 {
			for _, point := range analysis.KeyPoints {
				if point != "" {
					logger.Info("🔑 " + point)
				}
			}
		}
	} else {
		logger.WithFields(fields).Warn("Article processed without AI analysis")
		logger.Info("📄 Article: " + article.Title)
	}

	logger.Info("🔗 Read more: " + article.URL)
}
