# 🎌 Anime Api - Autonomous Sinhala Anime Blogger

An intelligent, autonomous system that discovers anime news, writes engaging content in authentic Sinhala, and publishes it automatically to social media. Powered by Google Gemini AI and designed to sound like a real Sri Lankan anime fan talking to friends.

## 🌟 Key Features

### 🤖 **Autonomous Operation**
- **Complete Automation**: Finds, writes, and publishes anime news autonomously
- **Smart Duplicate Detection**: Never posts the same news twice
- **Intelligent Content Filtering**: Only processes relevant anime/manga content

### 🇱🇰 **Authentic Sinhala Content**
- **Native Writing Style**: Uses casual, authentic Sinhala mixed with English (Singlish)
- **Cultural Context**: Writes like a real Sri Lankan anime fan
- **Engaging Tone**: Creates excitement and community engagement

### 📡 **Multi-Source RSS Integration**
- **Anime News Network**
- **Crunchyroll News**
- **MyAnimeList**
- **Otaku News**
- **Anime Hunch**

### 📱 **Social Media Publishing**
- **Telegram Bot Integration**
- **Twitter/X API Support** (extensible)
- **Automatic Status Tracking**

## 🏗️ Architecture Overview

The system follows an **autonomous agent pattern** with five specialized tools:

### **Tool 1: RSS News Fetcher** (`fetch_anime_news`)
```go
// Fetches latest anime news from multiple RSS feeds
articles, err := rssFetcher.FetchAnimeNews(ctx)
```

### **Tool 2: Duplicate Checker** (`check_if_posted_before`)
```go
// Checks if content was already published
isNew, err := duplicateChecker.CheckIfPostedBefore(articleLink)
```

### **Tool 3: AI Sinhala Writer** (`write_anime_post_in_my_style`)
```go
// Generates authentic Sinhala content using Gemini AI
sinhalaText, err := sinhalaWriter.WriteAnimePostInMyStyle(ctx, title, summary, link)
```

### **Tool 4: Social Media Publisher** (`publish_post`)
```go
// Publishes content to social media platforms
result, err := socialMediaPublisher.PublishPost(ctx, postText)
```

### **Tool 5: Publication Logger** (`log_as_published`)
```go
// Logs published content to prevent duplicates
err := duplicateChecker.LogAsPublished(articleLink, title)
```

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Google Gemini API key
- Telegram Bot (optional, for social media publishing)

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/AdhiDevX369/AclNews-go.git
cd AclNews-go
```

2. **Install dependencies**
```bash
go mod download
```

3. **Configure environment**
```bash
cp .env.example .env
# Edit .env with your API keys
```

4. **Set up API keys**

**Gemini API:**
- Visit [Google AI Studio](https://ai.google.dev/)
- Create a new API key
- Add to `.env` as `GEMINI_API_KEY`

**Telegram Bot (Optional):**
- Message [@BotFather](https://t.me/botfather) on Telegram
- Create a new bot and get the token
- Get your chat ID from [@userinfobot](https://t.me/userinfobot)
- Add both to `.env`

### Running the Application

**Run complete autonomous cycle:**
```bash
go run cmd/app/main.go
```

**CLI Commands:**
```bash
# Test all tools without posting
go run cmd/cli/main.go --test

# Check current status
go run cmd/cli/main.go --status

# Run one complete cycle
go run cmd/cli/main.go --run
```

## 🛠️ Development

### Build and Test
```bash
# Run tests
make test

# Build application
make build

# Run with hot reload
make dev

# Format code
make fmt

# Security scan
make security
```

### Docker Deployment
```bash
# Build Docker image
make docker-build

# Run with Docker Compose
docker-compose up -d
```

## 📊 Agent Workflow

The autonomous agent follows this reasoning sequence:

1. **🚀 Wake Up**: "Time to check for exciting anime news..."
2. **📡 Fetch News**: Uses RSS feeds to get latest articles
3. **🔍 Validate**: Checks each article for duplicates
4. **✍️ Write**: Creates authentic Sinhala content with AI
5. **📢 Publish**: Posts to social media platforms
6. **📋 Log**: Records publication to prevent duplicates
7. **😴 Sleep**: Waits for next scheduled execution

## 🎯 AI Persona

The Gemini AI acts as **"Anime Api"** (අනිමේ ඇපි) with these characteristics:

- **Personality**: Young, energetic Sri Lankan anime blogger
- **Audience**: Sri Lankan anime fans and enthusiasts
- **Style**: Casual, conversational Sinhala mixed with English
- **Goal**: Create excitement and community engagement
- **Expressions**: Uses authentic Sri Lankan slang and expressions

## 📁 Project Structure

```
anime-api/
├── cmd/
│   ├── app/           # Main application
│   └── cli/           # CLI tool
├── internal/
│   ├── config/        # Configuration management
│   ├── models/        # Data structures
│   └── services/      # Business logic
│       ├── rss_fetcher.go          # Tool 1: RSS fetching
│       ├── duplicate_checker.go    # Tool 2: Duplicate detection
│       ├── sinhala_writer.go       # Tool 3: AI content generation
│       ├── social_media_publisher.go # Tool 4: Social publishing
│       └── orchestrator.go         # Tool 5: Agent orchestrator
├── pkg/
│   ├── logger/        # Logging utilities
│   └── errors/        # Error handling
├── data/              # Persistent storage
└── docs/              # Documentation
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `GEMINI_API_KEY` | Google Gemini API key | ✅ Yes |
| `TELEGRAM_BOT_TOKEN` | Telegram bot token | 🔶 Optional |
| `TELEGRAM_CHAT_ID` | Telegram chat/channel ID | 🔶 Optional |
| `MAX_ARTICLES` | Max articles per cycle | ❌ No (default: 5) |
| `REQUEST_TIMEOUT` | API request timeout | ❌ No (default: 30s) |
| `ENVIRONMENT` | Environment (dev/prod) | ❌ No (default: dev) |

## 📈 Monitoring & Logging

### Status Monitoring
```bash
# Get current status
curl http://localhost:8080/health

# View recent publications
curl http://localhost:8080/status
```

### Logs
- **JSON Format**: Structured logging for production
- **Debug Mode**: Detailed tracing in development
- **Error Tracking**: Comprehensive error reporting

## 🔒 Security Features

- **API Key Validation**: Secure credential management
- **Rate Limiting**: Prevents API quota exhaustion
- **Input Sanitization**: Clean HTML and malicious content
- **Error Handling**: Graceful failure recovery

## 🚀 Deployment

### Production Deployment
```bash
# Build production image
make docker-build

# Deploy with environment
export GEMINI_API_KEY="your-key"
export TELEGRAM_BOT_TOKEN="your-token"
docker-compose -f docker-compose.prod.yml up -d
```

### Scheduled Execution
```bash
# Add to crontab for daily execution
0 9 * * * /path/to/anime-api --run
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Implement your changes
4. Add tests and documentation
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎯 Roadmap

- [ ] **Multi-platform Support**: Facebook, Instagram, TikTok integration
- [ ] **Advanced AI Features**: Image generation, video summaries
- [ ] **Community Features**: User interaction, polls, Q&A
- [ ] **Analytics Dashboard**: Performance metrics, engagement tracking
- [ ] **Mobile App**: iOS/Android companion app

---

**Made with ❤️ for the Sri Lankan anime community** 🇱🇰
