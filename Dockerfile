FROM golang:1.18-alpine as build-env
 
# Set environment variable
ENV APP_NAME readme-update-actions
ENV CMD_PATH main.go

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# install dependencies
RUN apk update && apk add --no-cache git

# Start app
ENTRYPOINT ["go", "run", "$APP_NAME/main.go"]
