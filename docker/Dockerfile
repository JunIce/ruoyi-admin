FROM nginx:1.18.0
MAINTAINER east.qrj@gmail.com

ADD dist/  /usr/share/nginx/html/
ADD nginx/default.conf /etc/nginx/conf.d/
ADD nginx/nginx.conf /etc/nginx/


ENV API_URL=https://vue.ruoyi.vip/prod-api/