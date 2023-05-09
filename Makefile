# build and save binary
build:
	go build -o dsim cmd/main.go

# build and run without saving binary
run:
	go run cmd/main.go

check: fmt vet lint test

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint ./...

test:
	go test ./...
