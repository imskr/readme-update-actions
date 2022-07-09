FROM golang:1.18-alpine as build-env
 
# Set environment variable
ENV APP_NAME readme-update-actions
ENV CMD_PATH main.go

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# install dependencies
RUN apk update && apk add --no-cache git

# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine:3.14

# install dependencies
RUN apk update && apk add --no-cache git

# Set environment variable
ENV APP_NAME readme-update-actions

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Start app
CMD ./$APP_NAME
