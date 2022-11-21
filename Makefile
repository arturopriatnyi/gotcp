all: lint build run

build:
	docker-compose build

run:
	docker-compose up --remove-orphans

lint:
	docker build . -f ./Dockerfile.lint -t linter
