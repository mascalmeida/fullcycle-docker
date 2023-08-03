# base image
FROM nginx

# remove .conf file from image
RUN rm /etc/nginx/conf.d/default.conf
# copy from local folder to docker image folder
COPY nginx.conf /etc/nginx/conf.d/default.conf

# create php file
#RUN mkdir /var/www/html -p && touch /var/www/html/index.php
