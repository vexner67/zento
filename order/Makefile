.PHONY: run build test fmt lint tidy proto check

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

proto:
	protoc \
		--proto_path=. \
		--go_out=. \
		--go_opt=module=github.com/vexner67/zento/order \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/vexner67/zento/order \
		api/order/v1/order.proto

check:
	$(MAKE) fmt
	$(MAKE) lint
	$(MAKE) test
