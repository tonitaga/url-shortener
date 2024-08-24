FROM golang:1.21.6-alpine3.19

RUN go version
ENV GOPATH=/

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o url-shortener ./cmd/app/main.go

EXPOSE 8000

CMD ["./url-shortener"]
