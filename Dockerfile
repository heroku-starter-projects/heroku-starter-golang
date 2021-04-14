FROM golang:1.16-alpine as builder

# https://stackoverflow.com/a/62265892/1217998
WORKDIR /app

# Download all the packages
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download && go mod verify

# Build the src
COPY . .
RUN go build -v -o server

# Runner stage
FROM alpine:latest

COPY --chown=0:0 --from=builder /app/server /app/server

EXPOSE 5001

CMD ["/app/server"]
