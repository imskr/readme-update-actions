FROM golang:1.18-alpine as build-env

WORKDIR /app
COPY . /app

# install dependencies
RUN apk update && apk add --no-cache git

# Start app
ENTRYPOINT ["go", "run", "/app/main.go"]
