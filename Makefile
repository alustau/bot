start:
	docker-compose up

bot:
	docker-compose up -d bot

mysql:
	docker-compose up -d mysql

build:
	docker build -t bot:latest .

test:
	go test ./...
