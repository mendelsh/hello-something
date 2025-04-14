# Build stage
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM scratch

COPY --from=builder /app/main /main

EXPOSE 8082

CMD ["/main"]
