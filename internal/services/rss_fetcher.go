package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go-test/internal/models"

	"github.com/mmcdole/gofeed"
)

// RSSFetcher handles fetching anime news from RSS feeds
type RSSFetcher struct {
	parser *gofeed.Parser
	feeds  []string
}

// NewRSSFetcher creates a new RSS fetcher instance
func NewRSSFetcher() *RSSFetcher {
	return &RSSFetcher{
		parser: gofeed.NewParser(),
		feeds: []string{
			"https://www.animenewsnetwork.com/all/rss.xml",
			"https://feeds.crunchyroll.com/news.rss",
			"https://myanimelist.net/rss/news.xml",
			"https://www.otakunews.com/feed/",
			"https://animehunch.com/feed/",
		},
	}
}

// FetchAnimeNews fetches latest anime news from RSS feeds
func (rf *RSSFetcher) FetchAnimeNews(ctx context.Context) ([]models.AnimeNews, error) {
	var allNews []models.AnimeNews

	for _, feedURL := range rf.feeds {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		news, err := rf.fetchFromFeed(ctx, feedURL)
		if err != nil {
			log.Printf("Error fetching from %s: %v", feedURL, err)
			continue // Continue with other feeds even if one fails
		}

		allNews = append(allNews, news...)
	}

	// Sort by published date (newest first)
	for i := 0; i < len(allNews)-1; i++ {
		for j := i + 1; j < len(allNews); j++ {
			if allNews[i].PublishedAt.Before(allNews[j].PublishedAt) {
				allNews[i], allNews[j] = allNews[j], allNews[i]
			}
		}
	}

	// Return top 15 most recent
	if len(allNews) > 15 {
		allNews = allNews[:15]
	}

	return allNews, nil
}

func (rf *RSSFetcher) fetchFromFeed(ctx context.Context, feedURL string) ([]models.AnimeNews, error) {
	feed, err := rf.parser.ParseURLWithContext(feedURL, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to parse feed %s: %w", feedURL, err)
	}

	var news []models.AnimeNews

	for _, item := range feed.Items {
		if item == nil {
			continue
		}

		// Filter for anime-related content
		if !rf.isAnimeRelated(item.Title, item.Description) {
			continue
		}

		publishedAt := time.Now()
		if item.PublishedParsed != nil {
			publishedAt = *item.PublishedParsed
		}

		description := ""
		if item.Description != "" {
			description = rf.cleanHTML(item.Description)
		}

		newsItem := models.AnimeNews{
			Title:       rf.cleanHTML(item.Title),
			Summary:     description,
			Link:        item.Link,
			Source:      rf.extractSourceName(feedURL),
			PublishedAt: publishedAt,
		}

		news = append(news, newsItem)
	}

	return news, nil
}

func (rf *RSSFetcher) isAnimeRelated(title, description string) bool {
	content := strings.ToLower(title + " " + description)

	// Check for anime-related keywords
	animeKeywords := []string{
		"anime", "manga", "otaku", "cosplay", "convention",
		"studio", "season", "episode", "character", "series",
		"crunchyroll", "funimation", "viz", "shonen", "seinen",
		"shoujo", "josei", "mecha", "isekai", "slice of life",
		"romance", "action", "adventure", "fantasy", "sci-fi",
		"dragon ball", "naruto", "one piece", "attack on titan",
		"demon slayer", "my hero academia", "jujutsu kaisen",
	}

	for _, keyword := range animeKeywords {
		if strings.Contains(content, keyword) {
			return true
		}
	}

	return false
}

func (rf *RSSFetcher) cleanHTML(content string) string {
	// Simple HTML tag removal
	content = strings.ReplaceAll(content, "<br>", " ")
	content = strings.ReplaceAll(content, "<br/>", " ")
	content = strings.ReplaceAll(content, "<br />", " ")
	content = strings.ReplaceAll(content, "</p>", " ")
	content = strings.ReplaceAll(content, "<p>", "")

	// Remove common HTML tags
	tags := []string{"<b>", "</b>", "<i>", "</i>", "<strong>", "</strong>",
		"<em>", "</em>", "<u>", "</u>", "<div>", "</div>", "<span>", "</span>"}

	for _, tag := range tags {
		content = strings.ReplaceAll(content, tag, "")
	}

	// Clean up extra whitespace
	content = strings.TrimSpace(content)
	words := strings.Fields(content)
	return strings.Join(words, " ")
}

func (rf *RSSFetcher) extractSourceName(feedURL string) string {
	if strings.Contains(feedURL, "animenewsnetwork") {
		return "Anime News Network"
	}
	if strings.Contains(feedURL, "crunchyroll") {
		return "Crunchyroll"
	}
	if strings.Contains(feedURL, "myanimelist") {
		return "MyAnimeList"
	}
	if strings.Contains(feedURL, "otakunews") {
		return "Otaku News"
	}
	if strings.Contains(feedURL, "animehunch") {
		return "Anime Hunch"
	}
	return "Unknown Source"
}
