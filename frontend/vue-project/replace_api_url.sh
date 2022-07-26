#!/usr/bin/env sh

find '/usr/share/nginx/html' -name '*.js' -exec sed -i -e 's,_API_BASE_URL,'"$API_BASE_URL"',g' {} \;
nginx -g "daemon off;"

find '/usr/share/nginx/html' -name '*.js' -exec sed -i -e 's,_API_PARSER_URL,'"$API_PARSER_URL"',g' {} \;
nginx -g "daemon off;"