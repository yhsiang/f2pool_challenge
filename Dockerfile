FROM golang:1.19.3-alpine AS build

WORKDIR /app

COPY . ./
RUN go mod download
COPY version.txt /app/restapi/handlers

RUN go build -o server ./cmd/interview-challenge-server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .
CMD ["/app/server", "--host", "0.0.0.0", "--port", "3000"]