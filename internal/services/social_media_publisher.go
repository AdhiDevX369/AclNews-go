package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SocialMediaPublisher handles publishing posts to social media platforms
type SocialMediaPublisher struct {
	telegramBotToken string
	telegramChatID   string
	httpClient       *http.Client
}

// NewSocialMediaPublisher creates a new social media publisher instance
func NewSocialMediaPublisher(telegramBotToken, telegramChatID string) *SocialMediaPublisher {
	return &SocialMediaPublisher{
		telegramBotToken: telegramBotToken,
		telegramChatID:   telegramChatID,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// TelegramMessage represents a Telegram API message
type TelegramMessage struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

// TelegramResponse represents the response from Telegram API
type TelegramResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

// PublishPost publishes a post to the configured social media platform
func (smp *SocialMediaPublisher) PublishPost(ctx context.Context, postText string) (string, error) {
	if smp.telegramBotToken != "" && smp.telegramChatID != "" {
		return smp.publishToTelegram(ctx, postText)
	}
	
	return "", fmt.Errorf("no social media platform configured")
}

// publishToTelegram publishes a post to Telegram
func (smp *SocialMediaPublisher) publishToTelegram(ctx context.Context, postText string) (string, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", smp.telegramBotToken)
	
	message := TelegramMessage{
		ChatID: smp.telegramChatID,
		Text:   postText,
	}
	
	jsonBody, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal Telegram message: %w", err)
	}
	
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create Telegram request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := smp.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send Telegram message: %w", err)
	}
	defer resp.Body.Close()
	
	var telegramResp TelegramResponse
	if err := json.NewDecoder(resp.Body).Decode(&telegramResp); err != nil {
		return "", fmt.Errorf("failed to decode Telegram response: %w", err)
	}
	
	if !telegramResp.OK {
		return "", fmt.Errorf("Telegram API error: %s (code: %d)", telegramResp.Description, telegramResp.ErrorCode)
	}
	
	return "Successfully posted to Telegram!", nil
}

// TestConnection tests the connection to the social media platform
func (smp *SocialMediaPublisher) TestConnection(ctx context.Context) error {
	if smp.telegramBotToken != "" && smp.telegramChatID != "" {
		return smp.testTelegramConnection(ctx)
	}
	
	return fmt.Errorf("no social media platform configured")
}

// testTelegramConnection tests the Telegram bot connection
func (smp *SocialMediaPublisher) testTelegramConnection(ctx context.Context) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", smp.telegramBotToken)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create test request: %w", err)
	}
	
	resp, err := smp.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to test Telegram connection: %w", err)
	}
	defer resp.Body.Close()
	
	var telegramResp TelegramResponse
	if err := json.NewDecoder(resp.Body).Decode(&telegramResp); err != nil {
		return fmt.Errorf("failed to decode test response: %w", err)
	}
	
	if !telegramResp.OK {
		return fmt.Errorf("Telegram bot test failed: %s", telegramResp.Description)
	}
	
	return nil
}
