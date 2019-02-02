FROM nginx:1.13-alpine

ADD ./dockerfiles/nginx/echo.conf.tpl /etc/nginx/conf.d/echo.conf.tpl
ADD ./dockerfiles/nginx/entrypoint.sh /var/opt/entrypoint.sh

ENTRYPOINT ["/var/opt/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
