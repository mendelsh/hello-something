# Build stage
FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod tidy

COPY app/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Final stage
FROM scratch

COPY --from=builder /app/main /main

EXPOSE 8082

CMD ["/main"]
