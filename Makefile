include .env

BINARY_NAME=main.go
PROJECT_NAME=$(shell basename "$(PWD)")

STDERR=/tmp/.$(PROJECTNAME)-stderr.txt
MAKE_FLAGS += --silent

build:
	go build -o ${BINARY_NAME}

run:
	docker-compose up -d
	go run ./${BINARY_NAME}
clean:
	go clean
	rm ${BINARY_NAME}
