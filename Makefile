check: fmt vet lint test

install: check
	go build -i -o $(GOPATH)/bin/dsim cmd/main.go

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint ./...

test:
	go test ./...