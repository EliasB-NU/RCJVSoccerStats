LABEL maintainer="Elias Braun"
# Build the Go application
FROM golang:1.25-alpine AS builder-go

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY backend/ ./backend/

RUN go build ./backend/main.go

# Build the Node.js application
FROM node:24.13.0-alpine AS builder-node

WORKDIR /app

COPY frontend/ ./

RUN npm install

RUN npm run build

# Final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates wget

WORKDIR /app

COPY --from=builder-go /app/main .

COPY --from=builder-node /app/dist ./web/dist

EXPOSE 3030

CMD ["./main"]