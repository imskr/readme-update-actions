FROM golang:1.18-alpine

RUN apk add -q --update \
    && apk add -q \
            bash \
            git \
            curl \
    && rm -rf /var/cache/apk/*

# Copy all the files from the host into the container
WORKDIR /src
COPY . .

# Enable Go modules
ENV GO111MODULE=on

# Compile the action
RUN go build -o /bin/action

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/action"]
