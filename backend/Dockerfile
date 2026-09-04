# Build stage
FROM golang:1.26.5-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN CGO_ENABLED=0 go build -o main main.go

# # Final stage — small runtime image
# FROM debian:bookworm-slim

# WORKDIR /app

# COPY --from=builder /app/main .

EXPOSE 8080
CMD go run main.go