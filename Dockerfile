FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache make git

RUN go mod download && go mod verify

RUN go build -o bin/scheduler

EXPOSE 8080

CMD ["./bin/scheduler"]