.PHONY: run build test fmt vet tidy

run:
	go run ./cmd/order

build:
	go build ./cmd/order

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

check:
	$(MAKE) fmt
	$(MAKE) lint
	$(MAKE) test
