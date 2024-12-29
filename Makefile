all: build

build:
	mkdir -p ./bin
	go run ./.github/cmd/main.go export-md --version=$(VERSION) > ./bin/go-dsa.md

test:
	go test --race -v -coverprofile=profile.cov ./...

clean:
	go clean
	rm -f ./bin/*

fmt:
	go fmt ./...
	gofmt -w -s -r 'interface{} -> any' .

deps:
	go mod tidy
	go mod download

vet:
	go vet ./...

ci: fmt vet test build clean

.PHONY: all build run test bench clean fmt deps update
