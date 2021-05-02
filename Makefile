test:
	go test ./...

install:
	go get

lint:
	golangci-lint run ./...

cover:
	mkdir -p tmp
	go test ./... -coverprofile=tmp/cover.out -coverpkg=./...
	go tool cover -html=tmp/cover.out -o tmp/cover.html
