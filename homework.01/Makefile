
TARGET ?= social-network

buf-build:
	buf generate -o pkg/grpc

buf-download:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0 \
		github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@v1.4.0 \
		github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@v1.4.0

golangci-lint:
	find -type f -name "*.go" | grep -v '.*\.pb\.go' | grep -v '\/[0-9a-z_]*.go' && echo "Files should be named in snake case" && exit 1 || echo "All files named in snake case"
	golangci-lint version
	golangci-lint run

buf-lint:
	buf lint

build:
	go build -o bin/$(TARGET) ./cmd

serve:
	go run ./cmd

fmt:
	gofumpt -w .
	gci write . --skip-generated -s standard -s default

tidy:
	go mod tidy

download:
	go mod download

modup: tidy
	go get -u ./...
	go mod tidy
