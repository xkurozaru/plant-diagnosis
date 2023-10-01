FROM node:18.17-slim

WORKDIR /app

COPY package.json yarn.lock /app/
RUN yarn install

COPY . /app
