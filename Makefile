
test:
	go test ./...

install:
	go get

lint:
	golangci-lint run ./...
