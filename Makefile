up:
	docker-compose -f docker/docker-compose.yml up --build -d

setup:
	cp .env.dist .env

compile:
	GOOS=linux GOARCH=386 go build -o main main.go

build: up setup compile

run:
	./main