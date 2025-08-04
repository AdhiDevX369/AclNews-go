package services

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-test/internal/models"
)

// DuplicateChecker manages tracking of published articles
type DuplicateChecker struct {
	logFilePath string
}

// NewDuplicateChecker creates a new duplicate checker instance
func NewDuplicateChecker() *DuplicateChecker {
	// Create data directory if it doesn't exist
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Warning: Could not create data directory: %v\n", err)
	}

	return &DuplicateChecker{
		logFilePath: filepath.Join(dataDir, "published_articles.txt"),
	}
}

// CheckIfPostedBefore checks if an article link has been posted before
func (dc *DuplicateChecker) CheckIfPostedBefore(articleLink string) (bool, error) {
	// If log file doesn't exist, article is definitely new
	if _, err := os.Stat(dc.logFilePath); os.IsNotExist(err) {
		return true, nil // true means it's NEW (not posted before)
	}

	file, err := os.Open(dc.logFilePath)
	if err != nil {
		return false, fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Each line format: "YYYY-MM-DD|LINK|TITLE"
		parts := strings.Split(line, "|")
		if len(parts) >= 2 && parts[1] == articleLink {
			return false, nil // false means it's OLD (already posted)
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error reading log file: %w", err)
	}

	return true, nil // true means it's NEW (not found in log)
}

// LogAsPublished logs an article as published
func (dc *DuplicateChecker) LogAsPublished(articleLink, title string) error {
	file, err := os.OpenFile(dc.logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file for writing: %w", err)
	}
	defer file.Close()

	// Format: "YYYY-MM-DD|LINK|TITLE"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s|%s|%s\n", timestamp, articleLink, title)

	if _, err := file.WriteString(logEntry); err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}

	return nil
}

// GetPublishedCount returns the number of published articles
func (dc *DuplicateChecker) GetPublishedCount() (int, error) {
	if _, err := os.Stat(dc.logFilePath); os.IsNotExist(err) {
		return 0, nil
	}

	file, err := os.Open(dc.logFilePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading log file: %w", err)
	}

	return count, nil
}

// GetRecentPublished returns recently published articles
func (dc *DuplicateChecker) GetRecentPublished(limit int) ([]models.PublishedArticle, error) {
	if _, err := os.Stat(dc.logFilePath); os.IsNotExist(err) {
		return []models.PublishedArticle{}, nil
	}

	file, err := os.Open(dc.logFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	var articles []models.PublishedArticle
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Parse line: "YYYY-MM-DD HH:MM:SS|LINK|TITLE"
		parts := strings.Split(line, "|")
		if len(parts) >= 3 {
			publishedAt, err := time.Parse("2006-01-02 15:04:05", parts[0])
			if err != nil {
				continue // Skip malformed entries
			}

			article := models.PublishedArticle{
				Link:        parts[1],
				Title:       parts[2],
				PublishedAt: publishedAt,
			}
			articles = append(articles, article)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %w", err)
	}

	// Return most recent articles (reverse order)
	start := len(articles) - limit
	if start < 0 {
		start = 0
	}

	recent := articles[start:]

	// Reverse to get newest first
	for i := 0; i < len(recent)/2; i++ {
		j := len(recent) - 1 - i
		recent[i], recent[j] = recent[j], recent[i]
	}

	return recent, nil
}
