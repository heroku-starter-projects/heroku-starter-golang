start:
	ENV=development go run main.go

build:
	go build -v

local-setup:
	cp ./env_sample ./.env

.PHONY: start start-dev build
