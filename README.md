# heroku-starter-golang (WORK IN PROGRESS)

A dockerized starter project which deploys an golang CRUD app to heroku.

Deployed here on heroku
<https://heroku-starter-golang.herokuapp.com>

## Installation

Make sure [Docker](https://www.docker.com/), [make](https://stackoverflow.com/a/11935185/1217998), [nodejs](https://nodejs.org/en/) (development only) and [golang](https://golang.org/doc/install) are installed

1. Run `docker-compose up` to start the postgres DB
2. Run `make local-setup` to install `nodemon` and setup environment
3. Run `make start` to run the migrations

## Environment variables

| Name        | Description                 |
| ----------- | --------------------------- |
| ENV         | `development`/`production`  |
| DB_USERNAME | Username of the postgres DB |
| DB_PASSWORD | Password of the postgres DB |
| DB_NAME     | Name of the postgres DB     |
| DB_HOST     | Host of the postgres DB     |

## CI Secrets

| Name            | Description                  |
| --------------- | ---------------------------- |
| HEROKU_API_KEY  | Heroku API key               |
| HEROKU_APP_NAME | Heroku app name              |
| HEROKU_EMAIL    | Your email address on Heroku |

## Migrations

This repo uses [GORM](https://gorm.io/docs/migration.html) for auto migration generation. The process is non-destructive.

## Running

Run `make start` to start in development mode.

Import the `./postman/*` into [Postman](https://www.postman.com/) to see the API docs
