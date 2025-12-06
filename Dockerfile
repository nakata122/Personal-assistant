FROM node:25-alpine AS client-builder
WORKDIR /client

COPY client/package.json .

RUN npm install

COPY client/ .

RUN npm run build

FROM golang:1.25 AS server-builder

WORKDIR /server

COPY server/go.mod server/go.sum ./

RUN go mod tidy

COPY --from=client-builder /client/dist ./dist

COPY server/ .

RUN go build -o bin/main cmd/main.go

RUN chmod +x /server/bin/main

CMD ["/server/bin/main"]
