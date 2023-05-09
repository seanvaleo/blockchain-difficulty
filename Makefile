build:
	go build -o dsim cmd/main.go

check: fmt vet lint test

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint ./...

test:
	go test ./...
