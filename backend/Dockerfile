FROM golang:1.20 as builder

COPY . /app/

WORKDIR /app

RUN go mod tidy && go build -o main .

FROM alpine:latest

COPY --from=builder --chmod=755 /app/main /app/main

WORKDIR /config

CMD ["/app/main"]