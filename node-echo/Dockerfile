#
# source: https://nodejs.org/en/docs/guides/nodejs-docker-webapp/
#
FROM node:8

WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install --only=production
COPY . .

EXPOSE 8001
CMD [ "npm", "start" ]