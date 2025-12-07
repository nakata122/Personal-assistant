#FRONTEND
FROM node:25-alpine AS client-builder
WORKDIR /client

COPY client/package.json .

RUN npm install

COPY client/ .

RUN npm run build

#BACKEND
FROM golang:1.25 AS server-builder

WORKDIR /server

COPY server/go.mod server/go.sum ./

RUN go mod tidy

COPY --from=client-builder /client/dist ./dist

COPY server/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

#FINAL IMAGE
FROM alpine:3.19

WORKDIR /app

COPY --from=server-builder /server/app .
COPY --from=server-builder /server/dist ./dist

EXPOSE 3000

CMD ["./app"]
