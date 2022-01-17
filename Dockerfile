FROM golang:1.17-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/cars-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/cars .

EXPOSE 8080

CMD ["./out/cars"]
