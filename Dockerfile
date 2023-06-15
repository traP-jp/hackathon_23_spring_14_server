FROM golang:1.20.5-alpine
ENV DOCKERIZE_VERSION v0.7.0
WORKDIR /github.com/traP-jp/hackathon_23_spring_14_server
RUN apk add --update --no-cache git \
    &&  wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY . .
RUN go mod download
RUN go build -o app