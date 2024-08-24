build:
	@go build -o bin/app main.go

run: build
	@./bin/app | jq
