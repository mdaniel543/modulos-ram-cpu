# Stage 1 - the build process

#FROM node:16-alpine AS node

#WORKDIR /app

#COPY . .

#RUN yarn install

#ENV VITE_BACKEND_URL=http://localhost:3000

#RUN yarn build

# Stage 2 - the production environment

FROM nginx:1.21.6-alpine

COPY nginx.conf /etc/nginx/nginx.conf

COPY ./dist/ /usr/share/nginx/html

# Language: dockerfile