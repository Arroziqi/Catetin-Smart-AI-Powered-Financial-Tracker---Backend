swagger:
	swag init -g cmd/docs/docs.go

build:
	go build -o main ./cmd/main.go

watch:
	air -c .air.toml

run:
	go run ./cmd/main.go