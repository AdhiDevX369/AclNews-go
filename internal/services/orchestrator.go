package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-test/internal/models"
)

// AnimeApiOrchestrator is the main orchestrator that acts as the autonomous Gemini agent
type AnimeApiOrchestrator struct {
	rssFetcher           *RSSFetcher
	duplicateChecker     *DuplicateChecker
	sinhalaWriter        *SinhalaWriter
	socialMediaPublisher *SocialMediaPublisher
	logger               *log.Logger
}

// NewAnimeApiOrchestrator creates a new orchestrator instance
func NewAnimeApiOrchestrator(
	rssFetcher *RSSFetcher,
	duplicateChecker *DuplicateChecker,
	sinhalaWriter *SinhalaWriter,
	socialMediaPublisher *SocialMediaPublisher,
	logger *log.Logger,
) *AnimeApiOrchestrator {
	return &AnimeApiOrchestrator{
		rssFetcher:           rssFetcher,
		duplicateChecker:     duplicateChecker,
		sinhalaWriter:        sinhalaWriter,
		socialMediaPublisher: socialMediaPublisher,
		logger:               logger,
	}
}

// ExecuteCycle runs one complete cycle of the autonomous agent
func (aao *AnimeApiOrchestrator) ExecuteCycle(ctx context.Context) error {
	aao.logger.Println("🚀 Anime Api awakening! Time to check for exciting anime news...")

	// Tool 1: Fetch anime news
	aao.logger.Println("📡 Tool 1: Fetching latest anime news from RSS feeds...")
	articles, err := aao.rssFetcher.FetchAnimeNews(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch anime news: %w", err)
	}

	if len(articles) == 0 {
		aao.logger.Println("❌ No anime articles found. Sleeping until next cycle...")
		return nil
	}

	aao.logger.Printf("✅ Found %d potential articles. Now checking for new content...", len(articles))

	// Tool 2: Find first new article
	var selectedArticle *models.AnimeNews
	for i, article := range articles {
		aao.logger.Printf("🔍 Checking article %d: %s", i+1, article.Title)

		isNew, err := aao.duplicateChecker.CheckIfPostedBefore(article.Link)
		if err != nil {
			aao.logger.Printf("⚠️  Error checking duplicate for article %d: %v", i+1, err)
			continue
		}

		if isNew {
			selectedArticle = &article
			aao.logger.Printf("🎉 Found NEW article: %s", article.Title)
			break
		} else {
			aao.logger.Printf("⏭️  Already posted: %s", article.Title)
		}
	}

	if selectedArticle == nil {
		aao.logger.Println("😴 All articles have been posted before. Nothing new to share today!")
		return nil
	}

	// Tool 3: Write Sinhala post
	aao.logger.Println("✍️  Tool 3: Writing exciting Sinhala post using AI...")
	sinhalaText, err := aao.sinhalaWriter.WriteAnimePostInMyStyle(
		ctx,
		selectedArticle.Title,
		selectedArticle.Summary,
		selectedArticle.Link,
	)
	if err != nil {
		return fmt.Errorf("failed to write Sinhala post: %w", err)
	}

	aao.logger.Println("📝 AI has crafted the perfect post! Here's what it wrote:")
	aao.logger.Printf("---\n%s\n---", sinhalaText)

	// Tool 4: Publish post
	aao.logger.Println("📢 Tool 4: Publishing to social media...")
	publishResult, err := aao.socialMediaPublisher.PublishPost(ctx, sinhalaText)
	if err != nil {
		return fmt.Errorf("failed to publish post: %w", err)
	}

	aao.logger.Printf("🎊 %s", publishResult)

	// Tool 5: Log as published
	aao.logger.Println("📋 Tool 5: Logging article as published...")
	err = aao.duplicateChecker.LogAsPublished(selectedArticle.Link, selectedArticle.Title)
	if err != nil {
		return fmt.Errorf("failed to log published article: %w", err)
	}

	aao.logger.Println("✅ Article logged successfully!")

	// Final success message
	aao.logger.Println("🎯 Mission accomplished! Anime Api has successfully:")
	aao.logger.Printf("   • Found new anime news: %s", selectedArticle.Title)
	aao.logger.Printf("   • Written engaging Sinhala content")
	aao.logger.Printf("   • Published to social media")
	aao.logger.Printf("   • Logged to prevent duplicates")
	aao.logger.Println("😴 Anime Api is now sleeping until the next scheduled run...")

	return nil
}

// GetStatus returns the current status of the orchestrator
func (aao *AnimeApiOrchestrator) GetStatus(ctx context.Context) (*models.OrchestratorStatus, error) {
	publishedCount, err := aao.duplicateChecker.GetPublishedCount()
	if err != nil {
		return nil, fmt.Errorf("failed to get published count: %w", err)
	}

	recentArticles, err := aao.duplicateChecker.GetRecentPublished(5)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent articles: %w", err)
	}

	// Test connections
	var connectionStatus []models.ServiceStatus

	// Test social media connection
	socialMediaErr := aao.socialMediaPublisher.TestConnection(ctx)
	connectionStatus = append(connectionStatus, models.ServiceStatus{
		Name:   "Social Media",
		Status: aao.getStatusString(socialMediaErr == nil),
		Error:  aao.getErrorString(socialMediaErr),
	})

	return &models.OrchestratorStatus{
		LastRun:         time.Now(),
		PublishedCount:  publishedCount,
		RecentArticles:  recentArticles,
		ServiceStatuses: connectionStatus,
	}, nil
}

func (aao *AnimeApiOrchestrator) getStatusString(isHealthy bool) string {
	if isHealthy {
		return "✅ Healthy"
	}
	return "❌ Error"
}

func (aao *AnimeApiOrchestrator) getErrorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

// TestAllTools tests all the tools in the orchestrator
func (aao *AnimeApiOrchestrator) TestAllTools(ctx context.Context) error {
	aao.logger.Println("🧪 Testing all tools...")

	// Test RSS Fetcher
	aao.logger.Println("Testing RSS Fetcher...")
	articles, err := aao.rssFetcher.FetchAnimeNews(ctx)
	if err != nil {
		return fmt.Errorf("RSS Fetcher test failed: %w", err)
	}
	aao.logger.Printf("✅ RSS Fetcher: Found %d articles", len(articles))

	// Test Duplicate Checker
	aao.logger.Println("Testing Duplicate Checker...")
	count, err := aao.duplicateChecker.GetPublishedCount()
	if err != nil {
		return fmt.Errorf("Duplicate Checker test failed: %w", err)
	}
	aao.logger.Printf("✅ Duplicate Checker: %d articles in log", count)

	// Test Social Media Publisher
	aao.logger.Println("Testing Social Media Publisher...")
	err = aao.socialMediaPublisher.TestConnection(ctx)
	if err != nil {
		return fmt.Errorf("Social Media Publisher test failed: %w", err)
	}
	aao.logger.Println("✅ Social Media Publisher: Connection successful")

	// Test Sinhala Writer (if we have articles)
	if len(articles) > 0 {
		aao.logger.Println("Testing Sinhala Writer...")
		_, err = aao.sinhalaWriter.WriteAnimePostInMyStyle(
			ctx,
			"Test Article",
			"This is a test summary for anime news testing.",
			"https://example.com/test",
		)
		if err != nil {
			return fmt.Errorf("Sinhala Writer test failed: %w", err)
		}
		aao.logger.Println("✅ Sinhala Writer: AI response generated successfully")
	}

	aao.logger.Println("🎉 All tools tested successfully!")
	return nil
}
