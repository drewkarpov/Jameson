FROM node:16-alpine as builder

COPY package.json package-lock.json webpack.mix.js ./
COPY src ./src

RUN npm install && mkdir ./public
RUN npm run prod

FROM nginx:alpine

COPY ./config/nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /public /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]