# Define a default target that runs when you just call 'make'
all: build run

# Define the build target for npm
build:
	@echo "Running npm build..."
	cd frontend/app && npm run build

# Define the run target for Go
run:
	@echo "Running Go project..."
	go run cmd/server/main.go