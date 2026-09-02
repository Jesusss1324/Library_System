FROM golang:1.26.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8081

CMD ["go", "run", "./cmd/api"]