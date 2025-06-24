.PHONY: build run test clean deps

all: deps build

# install dependencies
deps:
	go mod tidy
	go mod download

# build the application
build:
	go build -o bin/todo-list cmd/main.go

# run the application
run:
	go run cmd/main.go

# test
test:
	go test ./...

# clean build files
clean:
	rm -rf bin/
	rm -rf data/

# create environment configuration file
env:
	cp env.example .env

# run in development mode
dev: env run

help:
	@echo "Available targets:"
	@echo "  deps    - Install dependencies"
	@echo "  build   - Build the application"
	@echo "  run     - Run the application"
	@echo "  test    - Run tests"
	@echo "  clean   - Clean build files"
	@echo "  env     - Create .env file from example"
	@echo "  dev     - Run in development mode"
	@echo "  prod    - Build for production"
	@echo "  help    - Show this help message"