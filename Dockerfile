FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dextrace ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/dextrace .

COPY .env .

USER appuser

EXPOSE 8080

CMD ["./dextrace"]