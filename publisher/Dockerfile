# Build first
FROM golang:1.23.0 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o appexec

# Run lightweight
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/appexec .

ENTRYPOINT ["./appexec"]
