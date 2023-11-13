run: build
	@./bin/sunny

build:
	@go build -o bin/sunny cmd/main/main.go