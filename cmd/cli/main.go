package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"go-test/internal/config"
	"go-test/internal/services"
)

func main() {
	var (
		testTools  = flag.Bool("test", false, "Test all tools without posting")
		showStatus = flag.Bool("status", false, "Show current status")
		runCycle   = flag.Bool("run", false, "Run one complete cycle")
	)
	flag.Parse()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize services
	stdLogger := log.New(os.Stdout, "[ANIME-API-CLI] ", log.LstdFlags)

	rssFetcher := services.NewRSSFetcher()
	duplicateChecker := services.NewDuplicateChecker()
	sinhalaWriter := services.NewSinhalaWriter(cfg.GeminiAPIKey)
	socialMediaPublisher := services.NewSocialMediaPublisher(cfg.TelegramBotToken, cfg.TelegramChatID)

	orchestrator := services.NewAnimeApiOrchestrator(
		rssFetcher,
		duplicateChecker,
		sinhalaWriter,
		socialMediaPublisher,
		stdLogger,
	)

	ctx := context.Background()

	switch {
	case *testTools:
		fmt.Println("🧪 Testing all tools...")
		if err := orchestrator.TestAllTools(ctx); err != nil {
			log.Fatalf("Tool testing failed: %v", err)
		}
		fmt.Println("✅ All tools tested successfully!")

	case *showStatus:
		fmt.Println("📊 Getting status...")
		status, err := orchestrator.GetStatus(ctx)
		if err != nil {
			log.Fatalf("Failed to get status: %v", err)
		}

		fmt.Printf("📈 Status Report:\n")
		fmt.Printf("   Published Articles: %d\n", status.PublishedCount)
		fmt.Printf("   Last Run: %s\n", status.LastRun.Format("2006-01-02 15:04:05"))
		fmt.Printf("   Recent Articles: %d\n", len(status.RecentArticles))

		for _, svc := range status.ServiceStatuses {
			fmt.Printf("   %s: %s", svc.Name, svc.Status)
			if svc.Error != "" {
				fmt.Printf(" (%s)", svc.Error)
			}
			fmt.Println()
		}

	case *runCycle:
		fmt.Println("🎯 Running complete autonomous cycle...")
		if err := orchestrator.ExecuteCycle(ctx); err != nil {
			log.Fatalf("Cycle failed: %v", err)
		}
		fmt.Println("🎉 Cycle completed successfully!")

	default:
		fmt.Println("Anime Api CLI Tool")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  --test    : Test all tools without posting")
		fmt.Println("  --status  : Show current status")
		fmt.Println("  --run     : Run one complete cycle")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  go run cmd/cli/main.go --test")
		fmt.Println("  go run cmd/cli/main.go --status")
		fmt.Println("  go run cmd/cli/main.go --run")
	}
}
