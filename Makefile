default: all

all: clean test build

.PHONY: clean
clean: 
	rm -rf ./build/* && go mod tidy

.PHONY: download
	@go mod download

.PHONY: build
build: 
	@go build -o ./build/flight_search cmd/main.go

.PHONY: build-docker
build-docker:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/flight_search cmd/main.go

.PHONY: test
test: 
	@go test -v ./...

.PHONY: test-cover
test-cover:
	@go test -cover -covermode=atomic ./...

.PHONY: run-local
run-local:
	go run ./cmd/main.go