# Docker file to deploy to prod
FROM node:18

WORKDIR /usr/src/app

COPY . .

RUN apt-get update \
    && npm init -y \
    && npm install express -y

EXPOSE 5000

CMD ["node", "index.js"]