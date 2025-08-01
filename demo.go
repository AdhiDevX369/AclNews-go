package main

import (
	"fmt"
	"time"
)

// Demo version without external APIs for testing
func main() {
	fmt.Println("🍃 Anime News AI Assistant (Demo Mode)")
	fmt.Println("======================================")

	// Mock news data for demonstration
	mockNews := []struct {
		Title       string
		Description string
		URL         string
		PublishedAt time.Time
	}{
		{
			Title:       "Attack on Titan Final Movie Announced",
			Description: "Studio WIT announces final movie to conclude the epic series",
			URL:         "https://example.com/aot-movie",
			PublishedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			Title:       "One Piece Chapter 1090 Breaks Records",
			Description: "Latest chapter reveals shocking revelations about the void century",
			URL:         "https://example.com/op-1090",
			PublishedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			Title:       "Demon Slayer Season 4 Production Begins",
			Description: "Ufotable announces production start for highly anticipated season",
			URL:         "https://example.com/ds-s4",
			PublishedAt: time.Now().Add(-30 * time.Minute),
		},
	}

	fmt.Println("📰 Processing mock anime news...")
	fmt.Printf("🤖 Found %d articles\n\n", len(mockNews))

	for i, article := range mockNews {
		fmt.Printf("📄 Article %d:\n", i+1)
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Published: %s\n", article.PublishedAt.Format("2006-01-02 15:04"))
		fmt.Printf("Description: %s\n", article.Description)

		// Mock AI analysis
		analysis := generateMockAnalysis(article.Title, article.Description)
		fmt.Printf("🧠 AI Analysis: %s\n", analysis)

		fmt.Printf("🔗 Read more: %s\n", article.URL)
		fmt.Println("--------------------------------------------------")
	}

	fmt.Println("\n✅ Demo complete!")
	fmt.Println("\n🔧 To use real APIs:")
	fmt.Println("1. Get API keys (see README.md)")
	fmt.Println("2. Create .env file with your keys")
	fmt.Println("3. Run the main app: go run cmd/app/main.go")
}

func generateMockAnalysis(title, description string) string {
	analyses := map[string]string{
		"Attack on Titan": "Epic finale alert! 🔥 This movie is set to deliver the emotional conclusion AoT fans have been waiting for. Expect mind-blowing reveals and tissue-worthy moments!",
		"One Piece":       "Oda strikes again! 📖 Chapter 1090 is packed with lore that'll have theorists working overtime. The void century secrets are finally coming to light!",
		"Demon Slayer":    "Ufotable's animation magic continues! ✨ Season 4 means more breathtaking fight scenes and Tanjiro's journey isn't over yet. Get ready for visual perfection!",
	}

	for keyword, analysis := range analyses {
		if contains(title, keyword) {
			return analysis
		}
	}

	return "Another exciting development in the anime world! This news is sure to get fans talking and speculating about what's coming next. 🎌"
}

func contains(text, substring string) bool {
	return len(text) >= len(substring) &&
		text[:len(substring)] == substring ||
		(len(text) > len(substring) && text[len(text)-len(substring):] == substring) ||
		(len(text) > len(substring) && findInString(text, substring))
}

func findInString(text, substring string) bool {
	for i := 0; i <= len(text)-len(substring); i++ {
		if text[i:i+len(substring)] == substring {
			return true
		}
	}
	return false
}
