FROM node:25-alpine AS frontend
WORKDIR /client

COPY client/package.json .

RUN npm install

COPY client/ .

RUN npm run build

FROM golang:1.25 AS server

WORKDIR /server

RUN mv client/dist .

COPY server/go.mod server/go.sum ./
RUN go mod tidy

COPY server/ .

RUN go build -o bin/main cmd/main.go

RUN chmod +x /server/bin/main

CMD ["/server/bin/main"]
