.PHONY: build run up down psql

build:
	go build -o dist/eventlite

run: build
	./dist/eventlite --config config.env

up:
	docker-compose up

down:
	docker-compose down

psql:
	docker-compose run postgres psql -h postgres postgres
