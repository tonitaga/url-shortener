FROM golang:1.21.6-alpine3.19

RUN go version
ENV GOPATH=/

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY postgres-wait.sh .
RUN chmod +x postgres-wait.sh
RUN apk add postgresql-client

COPY . .
RUN go build -o url-shortener ./cmd/app/main.go

CMD ["./url-shortener"]
