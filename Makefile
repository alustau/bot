start:
	docker-compose up -d mysql

install:
	go install

build:
	go build -o bot main.go

docker-build:
	docker build -t bot:cgauge .

test:
	go test ./...
