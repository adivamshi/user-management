# syntax=docker/dockerfile:1

FROM golang:1.22.2

WORKDIR /app

COPY . .

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8081
CMD ["make", "build", "run"]
