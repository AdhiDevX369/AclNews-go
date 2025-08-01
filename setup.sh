#!/bin/bash

echo "🍃 Anime News AI Setup"
echo "====================="

# Check if .env exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp .env.example .env
    echo "✅ .env file created!"
    echo ""
    echo "🔑 Please edit .env file and add your API keys:"
    echo "   - GEMINI_API_KEY: Get from https://aistudio.google.com/app/apikey"
    echo "   - NEWS_API_KEY: Get from https://newsapi.org/register"
    echo ""
    echo "After adding your keys, run: go run cmd/app/main.go"
else
    echo "✅ .env file already exists"
fi

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

echo ""
echo "🚀 Setup complete!"
echo "Next steps:"
echo "1. Edit .env file with your API keys"
echo "2. Run: go run cmd/app/main.go"
