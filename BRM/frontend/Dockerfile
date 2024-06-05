#STAGE 1
FROM node:latest AS build
WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

#STAGE 2
FROM nginx:stable-alpine3.17-slim
WORKDIR /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY --from=build /usr/src/app/dist/brm/browser .
#overriding default nginx configuration
COPY nginx.conf /etc/nginx/conf.d/