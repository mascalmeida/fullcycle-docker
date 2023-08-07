# Docker file to deploy to prod
FROM node:18

WORKDIR /usr/src/app

RUN npm init -y && npm install express --save && npm install mysql --save

COPY . .

EXPOSE 3000 

CMD ["node", "index.js"]