package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// API Configuration
	GeminiAPIKey string
	NewsAPIKey   string

	// Application Configuration
	Port        string
	Environment string

	// Rate Limiting
	MaxArticles    int
	RequestTimeout time.Duration
	RetryAttempts  int
	RateLimitDelay time.Duration

	// Security
	AllowedOrigins []string
	MaxRequestSize int64

	// Logging
	LogLevel  string
	LogFormat string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (ignore error in production)
	_ = godotenv.Load()

	cfg := &Config{
		// Required API keys
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),
		NewsAPIKey:   getEnv("NEWS_API_KEY", ""),

		// Application settings
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),

		// Rate limiting defaults
		MaxArticles:    getEnvAsInt("MAX_ARTICLES", 5),
		RequestTimeout: getEnvAsDuration("REQUEST_TIMEOUT", "30s"),
		RetryAttempts:  getEnvAsInt("RETRY_ATTEMPTS", 3),
		RateLimitDelay: getEnvAsDuration("RATE_LIMIT_DELAY", "1s"),

		// Security defaults
		AllowedOrigins: []string{"*"},                                // Configure properly in production
		MaxRequestSize: getEnvAsInt64("MAX_REQUEST_SIZE", 1024*1024), // 1MB

		// Logging defaults
		LogLevel:  getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "json"),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

// validate ensures all required configuration is present
func (c *Config) validate() error {
	if c.GeminiAPIKey == "" {
		return fmt.Errorf("GEMINI_API_KEY is required")
	}

	if c.NewsAPIKey == "" {
		return fmt.Errorf("NEWS_API_KEY is required")
	}

	if c.MaxArticles <= 0 || c.MaxArticles > 100 {
		return fmt.Errorf("MAX_ARTICLES must be between 1 and 100")
	}

	return nil
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// Helper functions for environment variable parsing
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	if duration, err := time.ParseDuration(defaultValue); err == nil {
		return duration
	}
	return 30 * time.Second // fallback
}
