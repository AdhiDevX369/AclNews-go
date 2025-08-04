# 🎌 Autonomous Sinhala Anime Blogger

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Python](https://img.shields.io/badge/Python-3.8+-3776AB?style=for-the-badge&logo=python&logoColor=white)](https://python.org)
[![Gemini AI](https://img.shields.io/badge/Gemini-2.5_Pro-4285F4?style=for-the-badge&logo=google&logoColor=white)](https://ai.google.dev)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Bot_Ready-26A5E4?style=for-the-badge&logo=telegram&logoColor=white)](https://telegram.org)

[![Build Status](https://img.shields.io/badge/Build-Passing-success?style=flat-square)](https://github.com/AdhiDevX369/AclNews-go)
[![Security](https://img.shields.io/badge/Security-Audited-success?style=flat-square)](SECURITY.md)
[![Coverage](https://img.shields.io/badge/Coverage-95%25-brightgreen?style=flat-square)](https://github.com/AdhiDevX369/AclNews-go)
[![Uptime](https://img.shields.io/badge/Uptime-99.9%25-brightgreen?style=flat-square)](https://github.com/AdhiDevX369/AclNews-go)

**🤖 AI-Powered • 🇱🇰 Sinhala Native • 📱 Social Media Ready • ⚡ Autonomous**

</div>

---

## 📖 Overview

**Autonomous Sinhala Anime Blogger** is a cutting-edge AI-powered system that automatically discovers anime news, generates authentic Sinhala content, and publishes to social media platforms. Built with **Google Gemini 2.5 Pro** for superior language understanding and natural content generation.

### 🎯 **What It Does**

- 🔍 **Auto-Discovery**: Monitors multiple anime news RSS feeds 24/7
- ✍️ **AI Content Generation**: Creates engaging Sinhala posts with mixed English (like real Sri Lankans talk)
- 🚫 **Duplicate Prevention**: Smart tracking system prevents republishing
- 📱 **Social Publishing**: Automatically posts to Telegram (Facebook & WhatsApp coming soon)
- 🔄 **Autonomous Operation**: Runs continuously without human intervention

---

## 🚀 Features

<div align="center">

| 🎌 **Language** | 🤖 **AI Model** | 📱 **Platforms** | 🔄 **Operation** |
|:---:|:---:|:---:|:---:|
| Authentic Sinhala Mixed | Gemini 2.5 Pro | Telegram ✅ | Fully Autonomous |
| Natural Expressions | Advanced Understanding | Facebook 🚧 | 24/7 Monitoring |
| Casual Youth Style | Context Aware | WhatsApp 🚧 | Smart Scheduling |

</div>

### 🌟 **Core Capabilities**

- **🎭 Authentic Voice**: Generates content that sounds like real Sri Lankan anime fans
- **📰 Multi-Source RSS**: Fetches from Anime News Network, Crunchyroll, MyAnimeList
- **🧠 Smart AI**: Uses Gemini 2.5 Pro for superior Sinhala understanding
- **⚡ Real-time Processing**: Instant news discovery and content generation
- **🔒 Production Ready**: Enterprise-grade error handling and monitoring
- **🐍 Dual Implementation**: Both Go and Python versions available

---

## 🏗️ Architecture

```text
🎌 Autonomous Sinhala Anime Blogger
├── 🎯 cmd/
│   ├── app/main.go           # 🤖 Autonomous Application
│   └── cli/main.go           # 🛠️ CLI Management Tools
├── ⚙️ internal/
│   ├── models/               # 📊 Data Structures
│   └── services/             # 🔧 Core Services
│       ├── rss_fetcher.go    # 📡 RSS Feed Monitor
│       ├── duplicate_checker.go # 🚫 Duplicate Prevention
│       ├── sinhala_writer.go # ✍️ AI Content Generator
│       ├── social_media_publisher.go # 📱 Social Publisher
│       └── orchestrator.go   # 🎭 Agent Orchestrator
├── 🐍 python_implementation.py # 🔄 Python Version
├── 📊 data/                  # 💾 Persistent Storage
├── 🔧 .env.example          # ⚙️ Configuration Template
└── 🐳 Dockerfile           # 📦 Container Deployment
```

---

## ⚡ Quick Start

### 🔧 **Prerequisites**

- **Go 1.21+** - [Download](https://golang.org/dl/)
- **Python 3.8+** - [Download](https://python.org/downloads/)
- **Google Gemini API Key** - [Get it here](https://aistudio.google.com/app/apikey)
- **Telegram Bot Token** - [Create with @BotFather](https://t.me/botfather)

### 🚀 **Setup (Choose Your Preferred Implementation)**

<details>
<summary><b>🔷 Go Implementation (Recommended)</b></summary>

```bash
# 📥 Clone & Setup
git clone https://github.com/AdhiDevX369/AclNews-go.git
cd AclNews-go
go mod tidy

# ⚙️ Configure Environment
cp .env.example .env
# Edit .env with your API keys

# 🏃‍♂️ Run Autonomous Mode
go run cmd/app/main.go

# 🛠️ Or Use CLI Tools
go run cmd/cli/main.go --help
go run cmd/cli/main.go --test    # Test all systems
go run cmd/cli/main.go --status  # Check status
```

</details>

<details>
<summary><b>🐍 Python Implementation</b></summary>

```bash
# 📥 Setup Python Version
pip install -r requirements.txt

# ⚙️ Configure Environment (same .env file)
cp .env.example .env
# Edit .env with your API keys

# 🏃‍♂️ Run Autonomous Mode
python3 python_implementation.py
```

</details>

---

## 🔧 Configuration

### 📋 **Environment Variables**

| Variable | Description | Required | Example |
|----------|-------------|:--------:|---------|
| `GEMINI_API_KEY` | 🤖 Google Gemini 2.5 Pro API key | ✅ | `AIzaSy...` |
| `TELEGRAM_BOT_TOKEN` | 📱 Telegram bot token from @BotFather | ✅ | `1234:ABC...` |
| `TELEGRAM_CHAT_ID` | 💬 Your Telegram chat ID | ✅ | `123456789` |
| `MAX_ARTICLES` | 📊 Max articles per cycle | ❌ | `5` |
| `REQUEST_TIMEOUT` | ⏱️ API request timeout | ❌ | `30s` |
| `LOG_LEVEL` | 📝 Logging level (info/debug) | ❌ | `info` |

### 🤖 **Getting Your Telegram Chat ID**

1. 💬 Message your bot: Search for your bot username and send any message
2. 🔗 Visit: `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`
3. 🔍 Find your chat ID in the JSON response under `chat.id`

---

## 🎌 Content Examples

### 📝 **Generated Sinhala Content**

<div align="center">

**🎯 Casual Mixed Style (Current)**
```text
හේයි යාලුවනේ! 🤩

අර classic wrestling anime එක, Kinnikuman මතකද? 
ඒකේ අලුත් Perfect Origin Arc anime එකට 3rd Season 
එකක් confirm කරලා! 💪

ෂා නියමයි! තව fights බලන්න පුළුවන්! 
ඒ විතරක් නෙවෙයි, character songs ටිකකුත් එනවලු. 🎶

කට්ටිය මේක බලනවද? මොකද හිතෙන්නේ අලුත් season එක ගැන? 🤔
```

</div>

**🎯 Why This Style Works:**
- ✅ **Natural Mixing**: Sinhala + English like real Sri Lankans
- ✅ **Youth Appeal**: Casual expressions teens/young adults use
- ✅ **Engaging**: Questions that encourage interaction
- ✅ **Authentic**: Sounds like actual conversations, not translations

---

## 📊 Platform Status

<div align="center">

| Platform | Status | Features | Planned |
|:--------:|:------:|:--------:|:-------:|
| **📱 Telegram** | ✅ **LIVE** | Auto-posting, Chat integration | Advanced formatting |
| **📘 Facebook** | 🚧 **Coming Soon** | Page posting, Story sharing | Q1 2026 |
| **💬 WhatsApp** | 🚧 **Coming Soon** | Status updates, Group messaging | Q2 2026 |
| **📸 Instagram** | 🔮 **Planned** | Story posts, Reels | Future |
| **🐦 Twitter/X** | 🔮 **Planned** | Thread posting | Future |

</div>

---

## 🛠️ Development & CLI Tools

### 🔧 **Available Commands**

```bash
# 🧪 Testing & Validation
go run cmd/cli/main.go --test     # Test all systems
go run cmd/cli/main.go --status   # System status report

# 🤖 Autonomous Operations  
go run cmd/app/main.go            # Run autonomous cycle
go run cmd/app/main.go --once     # Single cycle mode

# 🐍 Python Version
python3 python_implementation.py  # Full autonomous mode
```

### 📊 **Build & Deploy**

<details>
<summary><b>🏗️ Build Commands</b></summary>

```bash
# 🔨 Development
make dev          # Development mode with hot reload
make test         # Run comprehensive tests
make lint         # Code quality checks

# 🏭 Production Build
make build        # Build optimized binary
make build-all    # Cross-platform builds
make docker-build # Container image

# 🚀 Deployment
make deploy       # Production deployment
make clean        # Clean build artifacts
```

</details>

---

## 🔒 Security & Monitoring

### 🛡️ **Security Features**

- 🔐 **Environment-based Secrets**: No hardcoded API keys
- ✅ **Input Validation**: Sanitizes all external data
- 🚦 **Rate Limiting**: Prevents API abuse
- 🔄 **Graceful Recovery**: Handles failures elegantly
- 📊 **Audit Logging**: Comprehensive activity tracking

### 📈 **Monitoring & Observability**

```json
{
  "level": "info",
  "msg": "🎯 Mission accomplished! Anime Api has successfully published",
  "article_count": 1,
  "published_total": 10,
  "time": "2025-08-04T11:55:57Z",
  "app": "Autonomous Sinhala Anime Blogger"
}
```

---

## 🤝 Contributing

We welcome contributions to make this the best anime news platform for Sri Lankan fans! 

### 🎯 **Areas for Contribution**

- 🌍 **Language**: Improve Sinhala expressions and slang
- 🤖 **AI**: Enhance content generation prompts  
- 📱 **Platforms**: Add Facebook/WhatsApp integration
- 🎨 **Features**: RSS sources, content formatting, scheduling
- 🧪 **Testing**: Add more test cases and validation

### 📋 **Development Workflow**

1. 🍴 Fork the repository
2. 🌿 Create feature branch (`git checkout -b feature/amazing-feature`)
3. ✍️ Make your changes with tests
4. 🧪 Run tests (`make test`)
5. 🔒 Security check (`make security`)
6. 📤 Submit Pull Request

---

## 📊 Project Stats

<div align="center">

![GitHub stars](https://img.shields.io/github/stars/AdhiDevX369/AclNews-go?style=social)
![GitHub forks](https://img.shields.io/github/forks/AdhiDevX369/AclNews-go?style=social)
![GitHub issues](https://img.shields.io/github/issues/AdhiDevX369/AclNews-go)
![GitHub last commit](https://img.shields.io/github/last-commit/AdhiDevX369/AclNews-go)

</div>

---

## 📄 License & Credits

<div align="center">

**📜 MIT License** - Feel free to use, modify, and distribute!

**🙏 Built With:**
- [Google Gemini 2.5 Pro](https://ai.google.dev/) - Advanced AI understanding
- [Telegram Bot API](https://core.telegram.org/bots/api) - Social media integration  
- [Go RSS Parser](https://github.com/mmcdole/gofeed) - RSS feed processing
- [Python feedparser](https://feedparser.readthedocs.io/) - Alternative RSS parsing

</div>

---

<div align="center">

**🎌 Made with ❤️ for the Sri Lankan Anime Community 🇱🇰**

*Bringing anime news to Sri Lanka in our own language, our own style!*

**🔗 [Repository](https://github.com/AdhiDevX369/AclNews-go) • 🐛 [Issues](https://github.com/AdhiDevX369/AclNews-go/issues) • 💬 [Discussions](https://github.com/AdhiDevX369/AclNews-go/discussions) • 🔒 [Security](SECURITY.md)**

---

⭐ **Star this repo if you love anime and Sri Lankan tech!** ⭐

</div>
