FROM node:19-alpine as socketservice

WORKDIR /socketservice
COPY /socketservice/package*.json /socketservice/
RUN npm install
COPY /socketservice /socketservice/

EXPOSE 3001

CMD ["npm", "run", "prod"]