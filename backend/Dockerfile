FROM golang:1.20 as builder

COPY . /app/

WORKDIR /app

RUN go mod tidy && go build -o /app/main

FROM ubuntu:devel

COPY --from=builder /app/main /app/main

WORKDIR "/config"

CMD ["bash","-c","/app/main"]
