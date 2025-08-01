# Anime News AI Makefile

# Variables
APP_NAME=anime-news-ai
VERSION=1.0.0
BUILD_DIR=bin
CMD_DIR=cmd/app

# Security and code quality
.PHONY: security
security: security-vuln security-static

.PHONY: security-vuln
security-vuln:
	@echo "🔒 Running vulnerability check..."
	@$(HOME)/go/bin/govulncheck ./...

.PHONY: security-static
security-static:
	@echo "🔍 Running static security analysis..."
	@$(HOME)/go/bin/staticcheck ./...

.PHONY: security-deps
security-deps:
	@echo "📦 Checking dependencies for known vulnerabilities..."
	@echo "External dependencies:"
	@go list -deps ./... | grep -v "go-test" | head -20rameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

.PHONY: all build clean test coverage lint fmt help install run dev

# Default target
all: clean fmt lint test build

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

# Build for multiple platforms
build-all: build-linux build-windows build-darwin

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux $(CMD_DIR)/main.go

build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows.exe $(CMD_DIR)/main.go

build-darwin:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin $(CMD_DIR)/main.go

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Lint the code
lint:
	@echo "Running linter..."
	$(GOVET) ./...

# Format the code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Install dependencies
install:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	$(GOCMD) run $(CMD_DIR)/main.go

# Run in development mode
dev: fmt lint
	@echo "Running in development mode..."
	LOG_LEVEL=debug $(GOCMD) run $(CMD_DIR)/main.go

# Setup environment
setup:
	@echo "Setting up environment..."
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo ".env file created from template"; \
		echo "Please edit .env file with your API keys"; \
	else \
		echo ".env file already exists"; \
	fi

# Security scan (requires gosec)
security:
	@if command -v gosec >/dev/null 2>&1; then \
		echo "Running security scan..."; \
		gosec ./...; \
	else \
		echo "gosec not installed. Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):$(VERSION) .
	docker tag $(APP_NAME):$(VERSION) $(APP_NAME):latest

# Docker run
docker-run:
	@echo "Running Docker container..."
	docker run --rm --env-file .env $(APP_NAME):latest

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Run clean, fmt, lint, test, and build"
	@echo "  build        - Build the application"
	@echo "  build-all    - Build for all platforms"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  coverage     - Run tests with coverage"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  install      - Install dependencies"
	@echo "  run          - Run the application"
	@echo "  dev          - Run in development mode"
	@echo "  setup        - Setup environment (.env file)"
	@echo "  security     - Run security scan"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  help         - Show this help message"
