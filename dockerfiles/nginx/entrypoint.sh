#!/bin/sh

envsubst '$$ECHO_SERVER $$ECHO_PORT' < /etc/nginx/conf.d/echo.conf.tpl > /etc/nginx/conf.d/echo.conf

exec "$@"
