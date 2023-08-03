# Docker file to deploy to prod
FROM node:18

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y wget

COPY . .

RUN npm init -y && npm install express -y

RUN npm install

EXPOSE 5000 

CMD ["node", "index.js"]