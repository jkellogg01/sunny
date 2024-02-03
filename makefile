run: build
	@bin/sunny

build:
	@go build -o bin/sunny main.go

clean:
	@$(RM) bin/*