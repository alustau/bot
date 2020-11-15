start:
	docker-compose up

build:
	docker build -t bot:latest .

test:
	go test ./...
