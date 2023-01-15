build:
	go build -o server main.go

run:
	go run main.go

tidy:
	go mod tidy

test:
	go test -v ./...

compile:
	GOOS=linux GOARCH=amd64 go build -o server main.go

compose-run:
	podman-compose -f docker-compose.dev.yaml up --build

watch:
	~/go/bin/reflex -s -r '\.go$$' make run