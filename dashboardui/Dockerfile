FROM node:19-alpine as dashboardui

WORKDIR /dashboardui
COPY /dashboardui/package*.json /dashboardui/
RUN npm install
COPY /dashboardui /dashboardui/
RUN npm run build

# production environment
FROM nginx:stable-alpine
COPY --from=dashboardui /dashboardui/out/ /usr/share/nginx/html
COPY /nginx/nginx.conf /etc/nginx/nginx.conf
COPY /nginx/sites/ /etc/nginx/sites

EXPOSE 80