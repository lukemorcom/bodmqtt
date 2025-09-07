# Build first
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o publisher .

# Run lightweight
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/publisher .

ENTRYPOINT ["./publisher"]
