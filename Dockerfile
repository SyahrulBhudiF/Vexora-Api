FROM golang:1.23 AS builder

WORKDIR /build

COPY go.mod go.sum config.yaml ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build cmd/vexora/vexora.go

FROM alpine:latest

# Copy the ca-certificates from the builder stage
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

# Optional: Copy host certificates (if required)
# COPY /etc/ssl/certs /etc/ssl/certs

RUN apk add htop

COPY --from=builder ["/build/vexora", "/build/config.yaml", "/"]

CMD ["/vexora", "-c", "/config.yaml", "serve"]