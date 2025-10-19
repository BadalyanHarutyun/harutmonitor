APP_NAME := harutmonitor
MAIN := main.go

.PHONY: all build run clean tidy

# Default target
all: build

# Build the Go binary
build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) $(MAIN)

# Run the application directly
run:
	@echo "Running $(APP_NAME)..."
	go run $(MAIN)

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Format and tidy dependencies
tidy:
	@echo "Tidying modules..."
	go fmt ./...
	go mod tidy
