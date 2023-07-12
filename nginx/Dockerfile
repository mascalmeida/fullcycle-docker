# base image
FROM nginx:latest
# copy from local folder to docker image folder
COPY html /usr/share/nginx/html
# defining the entrypoit: it is fixed
ENTRYPOINT [ "/docker-entrypoint.sh" ]
# defining the command: it can be changed
CMD ["nginx", "-g", "daemon off;"]