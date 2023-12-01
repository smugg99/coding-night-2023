FROM nginx:mainline-alpine AS base

FROM node:18-alpine AS build
ARG SERVER_URL
ENV VITE_SERVER_URL $VITE_SERVER_URL
WORKDIR /app/
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY . .
RUN npm run build 

FROM base AS final
COPY --from=build /app/dist /usr/share/nginx/html
