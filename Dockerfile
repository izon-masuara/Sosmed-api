FROM golang:1.18.1-buster
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
CMD go run main.go