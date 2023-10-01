FROM node:18.17-slim

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install

COPY . .
