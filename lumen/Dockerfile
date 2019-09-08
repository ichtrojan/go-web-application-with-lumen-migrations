FROM php:7.2-fpm

RUN rm /etc/apt/preferences.d/no-debian-php

RUN apt-get update && apt-get install -y libmcrypt-dev \
    libxml2-dev \
    libpq-dev \
    php-soap \
    php-mysql \
    php-pgsql \
    && apt-get clean -y \
    && docker-php-ext-install pdo pdo_pgsql

RUN pecl install mcrypt-1.0.1 \
	&& docker-php-ext-enable mcrypt

WORKDIR /var/www

COPY . /var/www

CMD php -S 0.0.0.0:9000 -t public

EXPOSE 9000


