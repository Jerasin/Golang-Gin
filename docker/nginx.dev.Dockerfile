FROM nginx:1.21-alpine
RUN rm -f /etc/nginx/conf.d/*
COPY ./nginx/dev/ /etc/nginx/conf.d/