# --------
# frontend
# --------

FROM node:current-alpine3.22 AS frontend

WORKDIR /app/frontend

# install requirements
COPY ./frontend/package.json .
RUN npm install

# build
COPY ./frontend .
RUN npm run build

# -----
# nginx
# -----

FROM nginx:stable-alpine3.23-perl AS nginx

WORKDIR /app/frontend

RUN chown nginx:nginx /var/cache/nginx/

# main nginx conf
COPY ./web-server/nginx.conf /etc/nginx/nginx.conf
# included nginx confs
COPY ./web-server/http_site.conf /etc/nginx/conf.d/http_site.conf
COPY ./web-server/https_site.conf /etc/nginx/conf.d/https_site.conf

# copy built frontend files
COPY --from=frontend /app/frontend/public ./public
COPY --from=frontend /app/frontend/dist ./dist
