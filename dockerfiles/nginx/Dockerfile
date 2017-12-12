FROM nginx:1.13-alpine

ADD echo.conf.tpl /etc/nginx/conf.d/echo.conf.tpl
ADD entrypoint.sh /var/opt/entrypoint.sh

ENTRYPOINT ["/var/opt/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
