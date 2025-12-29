FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o tg_helper_bot ./cmd -token '8565782509:AAF5gy1a4bsFvS0_AV5KTH1Y6smTEheJ-xY'

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["./app"]