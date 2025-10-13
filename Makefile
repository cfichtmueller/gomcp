# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOLINT=golangci-lint

# Project parameters
PROJECT_NAME=gomcp
MODULE_NAME=github.com/cfichtmueller/gomcp
BUILD_DIR=build
EXAMPLES_DIR=examples

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)"

.PHONY: all build clean test test-verbose test-coverage fmt vet lint deps tidy examples help

# Default target
all: clean deps fmt vet lint test build

# Build the library (no binary output for library)
build:
	@echo "Building $(PROJECT_NAME) library..."
	@mkdir -p $(BUILD_DIR)
	@$(GOCMD) list -f '{{.Dir}}' ./... | head -1 > /dev/null
	@echo "Library build completed successfully"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -f $(EXAMPLES_DIR)/*/__debug_bin*
	@echo "Clean completed"

# Run tests
test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

# Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	@$(GOTEST) -v -race ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@$(GOTEST) -v -race -coverprofile=coverage.out ./...
	@$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
fmt:
	@echo "Formatting code..."
	@$(GOFMT) ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@$(GOVET) ./...

# Run linter (requires golangci-lint to be installed)
lint:
	@echo "Running linter..."
	@if command -v $(GOLINT) >/dev/null 2>&1; then \
		$(GOLINT) run ./...; \
	else \
		echo "golangci-lint not found. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		echo "Skipping linting..."; \
	fi

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@$(GOMOD) download

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@$(GOMOD) tidy

# Build examples
examples: build
	@echo "Building examples..."
	@for example in $(EXAMPLES_DIR)/*; do \
		if [ -d "$$example" ] && [ -f "$$example/main.go" ]; then \
			echo "Building $$example..."; \
			$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$$(basename $$example) $$example/main.go; \
		fi; \
	done
	@echo "Examples built successfully"

# Run hello server example
run-hello: examples
	@echo "Running hello server example..."
	@$(BUILD_DIR)/hello_server

# Install the library (for local development)
install:
	@echo "Installing $(PROJECT_NAME)..."
	@$(GOCMD) install ./...

# Generate documentation
docs:
	@echo "Generating documentation..."
	@$(GOCMD) doc -all ./... > $(BUILD_DIR)/docs.txt
	@echo "Documentation generated: $(BUILD_DIR)/docs.txt"

# Check for security vulnerabilities
security:
	@echo "Checking for security vulnerabilities..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not found. Install it with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
		echo "Skipping security check..."; \
	fi

# Benchmark tests
benchmark:
	@echo "Running benchmarks..."
	@$(GOTEST) -bench=. -benchmem ./...

# Check if all dependencies are available
check-deps:
	@echo "Checking dependencies..."
	@$(GOCMD) mod verify
	@echo "Dependencies verified"

# Create a release build
release: clean deps fmt vet lint test
	@echo "Creating release build..."
	@mkdir -p $(BUILD_DIR)/release
	@echo "Release build completed"

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	@$(GOMOD) download
	@$(GOMOD) tidy
	@echo "Development setup completed"

# Show help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, deps, fmt, vet, lint, test, and build"
	@echo "  build        - Build the library"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  test-verbose - Run tests with verbose output"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  lint         - Run linter (requires golangci-lint)"
	@echo "  deps         - Download dependencies"
	@echo "  tidy         - Tidy dependencies"
	@echo "  examples     - Build all examples"
	@echo "  run-hello    - Run hello server example"
	@echo "  install      - Install the library"
	@echo "  docs         - Generate documentation"
	@echo "  security     - Check for security vulnerabilities"
	@echo "  benchmark    - Run benchmark tests"
	@echo "  check-deps   - Verify dependencies"
	@echo "  release      - Create release build"
	@echo "  dev-setup    - Setup development environment"
	@echo "  help         - Show this help message"
