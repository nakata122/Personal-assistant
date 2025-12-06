FROM node:25-alpine AS frontend
WORKDIR /app

COPY client/package.json .

RUN npm install

RUN npm i -g serve

COPY client/ .

RUN npm run build

EXPOSE 3000

CMD [ "serve", "-s", "dist" ]