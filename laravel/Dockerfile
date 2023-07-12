# Base image
FROM php:7.4-cli
# Set directory
WORKDIR /var/wwww
# Install dependecies
RUN apt-get update && \
    apt-get install libzip-dev -y && \
    docker-php-ext-install zip
# Install composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');"
# Install laravel
RUN php composer.phar create-project --prefer-dist laravel/laravel mylaravel
# Defining entrypoint
ENTRYPOINT [ "php", "mylaravel/artisan", "serve" ]
# Defining command
CMD [ "--host=0.0.0.0" ]