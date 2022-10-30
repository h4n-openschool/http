deps:
	@go mod download

test: deps
	@go test ./...

example: deps
	@go run example/main.go
