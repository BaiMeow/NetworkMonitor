FROM node:18 as frontend

COPY frontend /app

WORKDIR /app

RUN npm install 

RUN npm run build

FROM golang:1.20 as builder

COPY backend /app

WORKDIR /app

RUN rm -rf /app/static

COPY --from=frontend /backend/static /app/static

RUN go mod tidy && go build -o /app/main

FROM ubuntu:devel

RUN apt update &&  \
    apt install -y ca-certificates &&  \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/main

WORKDIR "/config"

CMD ["bash","-c","/app/main"]

EXPOSE 8787