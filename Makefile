deps:
	@go mod download

test: deps
	@go test ./...
