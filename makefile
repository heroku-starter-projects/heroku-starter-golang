start:
	ENV=development nodemon -L --exec go run main.go --signal SIGTERM

build:
	go build -v

local-setup:
	sudo npm i -g nodemon
	cp ./env_sample ./.env

.PHONY: start start-dev build
