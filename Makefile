.PHONY: build clean deploy

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/gservice main.go

clean:
	rm -rf ./bin

test: build
	go test ./...

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down
