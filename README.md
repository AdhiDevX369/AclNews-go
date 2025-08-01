# 🍃 Anime News AI

A professional, secure, and scalable AI-powered anime news analysis application built with Go.

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Security](https://img.shields.io/badge/security-audited-green.svg)](SECURITY.md)

## 🚀 Features

- **Real-time News Fetching**: Automatically fetches latest anime and manga news from NewsAPI
- **AI-Powered Analysis**: Uses Google Gemini AI for intelligent content analysis and summarization
- **Professional Architecture**: Clean, modular, and maintainable codebase following Go best practices
- **Security-First Design**: Secure API key management, input validation, and error handling
- **Structured Logging**: JSON-formatted logs with configurable levels for monitoring
- **Graceful Shutdown**: Proper resource cleanup and signal handling
- **Rate Limiting**: Respects API limits with built-in retry logic and exponential backoff
- **Environment Configuration**: Flexible environment-based configuration management
- **Production Ready**: Docker support, health checks, and monitoring capabilities

## 🏗️ Architecture

```text
anime-news-ai/
├── cmd/app/                 # Application entry point
│   └── main.go             # Main application
├── internal/                # Private application code
│   ├── config/             # Configuration management
│   ├── models/             # Data models & types
│   └── services/           # Business logic services
│       ├── gemini.go       # Google Gemini AI service
│       └── news.go         # NewsAPI service
├── pkg/                    # Reusable packages
│   ├── errors/             # Error handling utilities
│   └── logger/             # Structured logging
├── .env.example            # Environment configuration template
├── Dockerfile              # Container deployment
├── Makefile               # Build automation
├── SECURITY.md            # Security documentation
└── README.md              # Project documentation
```

## 📋 Prerequisites

- **Go 1.21+** - [Download](https://golang.org/dl/)
- **Google Gemini AI API Key** - [Get it here](https://aistudio.google.com/app/apikey)
- **NewsAPI Key** - [Get it here](https://newsapi.org/register)

## ⚡ Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/AdhiDevX369/AclNews-go.git
cd AclNews-go
go mod tidy
```

### 2. Configure Environment

```bash
cp .env.example .env
# Edit .env with your actual API keys
```

### 3. Run the Application

```bash
# Using Go directly
go run cmd/app/main.go

# Or using Make
make run

# Or build and run
make build
./bin/anime-news-ai
```

## 🔧 Configuration

The application uses environment variables for configuration. Copy `.env.example` to `.env` and configure as needed.

### Required Configuration

| Variable | Description | Example |
|----------|-------------|---------|
| `GEMINI_API_KEY` | Google Gemini AI API key | `AIzaSy...` |
| `NEWS_API_KEY` | NewsAPI key for fetching news | `abc123...` |

### Optional Configuration

| Variable | Description | Default |
|----------|-------------|---------|
| `ENVIRONMENT` | Runtime environment (development/production) | `development` |
| `PORT` | Server port (if applicable) | `8080` |
| `MAX_ARTICLES` | Maximum articles to process per run | `5` |
| `LOG_LEVEL` | Logging level (debug/info/warn/error) | `info` |
| `LOG_FORMAT` | Log format (json/text) | `json` |
| `REQUEST_TIMEOUT` | API request timeout | `30s` |
| `RETRY_ATTEMPTS` | Number of retry attempts for failed requests | `3` |
| `RATE_LIMIT_DELAY` | Delay between API calls | `1s` |

## 🔒 Security Features

- **Environment-based Configuration**: No hardcoded secrets or API keys
- **Input Validation**: Validates and sanitizes all external inputs
- **Rate Limiting**: Built-in protection against API abuse
- **Secure Error Handling**: Error messages without sensitive information leakage
- **Request Timeouts**: Prevents resource exhaustion
- **Graceful Shutdown**: Proper cleanup on termination signals
- **Security Audit**: Regular security scanning with `govulncheck` and `staticcheck`

## 📊 Monitoring & Observability

### Structured Logging

The application uses structured JSON logging with configurable levels:

```json
{
  "level": "info",
  "msg": "Articles fetched successfully",
  "count": 5,
  "time": "2025-08-01T14:30:00Z",
  "app": "Anime News AI",
  "version": "1.0.0"
}
```

### Health Monitoring

- Service validation on startup
- API key verification
- Graceful error handling with retry logic
- Context-based cancellation
- Resource cleanup monitoring

## 🧪 Development

### Available Make Commands

```bash
# Development
make dev          # Run in development mode
make fmt          # Format code
make lint         # Run linter
make test         # Run tests
make coverage     # Generate test coverage

# Security
make security     # Run all security tests
make security-vuln # Run vulnerability scan
make security-static # Run static analysis

# Build & Deploy
make build        # Build application
make build-all    # Build for all platforms
make docker-build # Build Docker image
make docker-run   # Run Docker container
make clean        # Clean build artifacts
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run specific package tests
go test ./internal/services/...
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Security scan
make security
```

## 🚀 Production Deployment

### Environment Setup

```bash
ENVIRONMENT=production
LOG_LEVEL=info
LOG_FORMAT=json
MAX_ARTICLES=10
REQUEST_TIMEOUT=60s
RETRY_ATTEMPTS=5
```

### Docker Deployment

```bash
# Build Docker image
make docker-build

# Run with Docker
make docker-run

# Or manually
docker build -t anime-news-ai .
docker run --env-file .env anime-news-ai
```

### Health Checks

The application includes built-in health monitoring:

- API key validation on startup
- Service connectivity checks
- Graceful error recovery
- Resource usage monitoring

## 📖 API Integration

### NewsAPI Integration

The application fetches anime-related news using:

- **Query**: "anime OR manga OR otaku"
- **Language**: English
- **Sort**: Published date (newest first)
- **Rate Limiting**: Respects API quotas

### Google Gemini AI Analysis

Gemini AI provides structured analysis including:

- **Summary**: Concise article summary (2-3 sentences)
- **Key Points**: Important highlights (3-5 points)
- **Sentiment**: POSITIVE/NEUTRAL/NEGATIVE classification
- **Relevance**: Anime/manga relevance scoring

### Example Output

```json
{
  "article_title": "New Anime Series Announced",
  "ai_summary": "A new anime adaptation has been announced...",
  "ai_sentiment": "POSITIVE",
  "ai_key_points": [
    "Studio announced new series",
    "Based on popular manga",
    "Release scheduled for 2025"
  ]
}
```

## 🐛 Troubleshooting

### Common Issues

#### Configuration Errors

```text
Error: "GEMINI_API_KEY is required"
Solution: Create .env file with valid API keys
```

#### API Connection Issues

```text
Error: "News API request failed"
Solution: Check API key validity and quota limits
```

#### Service Validation Failed

```text
Error: "Service validation failed"
Solution: Verify both API keys are valid and have quota
```

### Debug Mode

Enable debug logging for detailed information:

```bash
LOG_LEVEL=debug make run
```

### Performance Tuning

```bash
# Increase article processing
MAX_ARTICLES=20

# Adjust timeouts for slow networks
REQUEST_TIMEOUT=60s

# Modify rate limiting
RATE_LIMIT_DELAY=2s
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`make test`)
5. Run security checks (`make security`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Development Guidelines

- Follow Go conventions and best practices
- Write tests for new functionality
- Update documentation as needed
- Run security scans before submitting
- Use structured logging for observability

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Google Gemini AI](https://ai.google.dev/) for AI analysis capabilities
- [NewsAPI](https://newsapi.org/) for comprehensive news data
- [Logrus](https://github.com/sirupsen/logrus) for structured logging
- [Godotenv](https://github.com/joho/godotenv) for environment management

## 🔗 Links

- **Repository**: [https://github.com/AdhiDevX369/AclNews-go](https://github.com/AdhiDevX369/AclNews-go)
- **Security Report**: [SECURITY.md](SECURITY.md)
- **Issues**: [GitHub Issues](https://github.com/AdhiDevX369/AclNews-go/issues)

---

Built with ❤️ for the anime community
