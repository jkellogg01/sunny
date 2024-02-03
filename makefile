run: build
<<<<<<< HEAD
	@bin/sunny

build:
	@go build -o bin/sunny main.go

clean:
	@$(RM) bin/*
=======
	@./bin/sunny

build:
	@go build -o bin/sunny cmd/main/main.go
>>>>>>> 032b1ec93653549021d4b547cca7b0f10038e462
