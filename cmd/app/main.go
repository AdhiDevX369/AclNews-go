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
	"go-test/internal/services"
	"go-test/pkg/logger"
)

const appVersion = "v1.0.0"

func main() {
	// Initialize logger
	appLogger := logger.New("info", "json")
	appLogger.Info("Starting Anime Api - Autonomous Sinhala Anime Blogger", map[string]interface{}{
		"version": appVersion,
		"time":    time.Now().Format("Monday, January 2, 2006 at 3:04 PM"),
	})

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		appLogger.Fatal("Failed to load configuration", map[string]interface{}{
			"error": err.Error(),
		})
	}

	appLogger.Info("Configuration loaded successfully", map[string]interface{}{
		"environment": cfg.Environment,
		"port":        cfg.Port,
	})

	// Initialize services
	services, err := initializeServices(cfg, appLogger)
	if err != nil {
		appLogger.Fatal("Failed to initialize services", map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Test all tools before starting
	appLogger.Info("Testing all tools before starting the cycle...")
	if err := services.orchestrator.TestAllTools(ctx); err != nil {
		appLogger.Error("Tool testing failed", map[string]interface{}{
			"error": err.Error(),
		})
		// Continue anyway, some tools might work
	}

	// Run one cycle
	appLogger.Info("🎯 Starting Anime Api autonomous cycle...")
	if err := services.orchestrator.ExecuteCycle(ctx); err != nil {
		appLogger.Error("Autonomous cycle failed", map[string]interface{}{
			"error": err.Error(),
		})
		os.Exit(1)
	}

	// Get status report
	status, err := services.orchestrator.GetStatus(ctx)
	if err != nil {
		appLogger.Error("Failed to get status", map[string]interface{}{
			"error": err.Error(),
		})
	} else {
		appLogger.Info("Final status report", map[string]interface{}{
			"published_count":  status.PublishedCount,
			"recent_articles":  len(status.RecentArticles),
			"last_run":        status.LastRun.Format("2006-01-02 15:04:05"),
		})
	}

	appLogger.Info("Anime Api cycle completed successfully! 🎉")
}

// ServiceContainer holds all initialized services
type ServiceContainer struct {
	orchestrator *services.AnimeApiOrchestrator
}

// initializeServices initializes all required services
func initializeServices(cfg *config.Config, appLogger *logger.Logger) (*ServiceContainer, error) {
	stdLogger := log.New(os.Stdout, "[ANIME-API] ", log.LstdFlags)

	// Initialize RSS Fetcher
	rssFetcher := services.NewRSSFetcher()

	// Initialize Duplicate Checker
	duplicateChecker := services.NewDuplicateChecker()

	// Initialize Sinhala Writer
	sinhalaWriter := services.NewSinhalaWriter(cfg.GeminiAPIKey)

	// Initialize Social Media Publisher
	socialMediaPublisher := services.NewSocialMediaPublisher(
		cfg.TelegramBotToken,
		cfg.TelegramChatID,
	)

	// Initialize Orchestrator
	orchestrator := services.NewAnimeApiOrchestrator(
		rssFetcher,
		duplicateChecker,
		sinhalaWriter,
		socialMediaPublisher,
		stdLogger,
	)

	return &ServiceContainer{
		orchestrator: orchestrator,
	}, nil
}

// setupGracefulShutdown sets up graceful shutdown handling
func setupGracefulShutdown(cancel context.CancelFunc, appLogger *logger.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		appLogger.Info("Received shutdown signal")
		cancel()
	}()
}

// displayWelcomeBanner displays the application welcome banner
func displayWelcomeBanner() {
	banner := `
╔══════════════════════════════════════════════════════════════╗
║                     🎌 ANIME API 🎌                          ║
║            Autonomous Sinhala Anime News Blogger            ║
║                                                              ║
║  🤖 Powered by Google Gemini AI                             ║
║  📡 RSS Feed Integration                                     ║
║  🇱🇰 Authentic Sinhala Content Generation                   ║
║  📱 Social Media Publishing                                  ║
╚══════════════════════════════════════════════════════════════╝
`
	fmt.Print(banner)
}
