FROM ubuntu:20.04

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://dl.google.com/go/go1.18.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.18.linux-amd64.tar.gz
RUN rm /tmp/go1.18.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go version

WORKDIR /app
COPY . /app
RUN chmod -R 777 /app

# Start app
ENTRYPOINT ["go"]
CMD ["run", "/app/main.go"]
