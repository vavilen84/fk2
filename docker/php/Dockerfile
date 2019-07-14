FROM php:7.3-fpm

ENV PHPREDIS_VERSION 3.1.6

# Mcrypt
RUN apt-get update && apt-get install -y libmcrypt-dev \
    && pecl install mcrypt-1.0.2 \
    && docker-php-ext-enable mcrypt

# GD
RUN apt-get update && apt-get install -y libpng-dev libwebp-dev libfreetype6-dev libjpeg62-turbo-dev libxpm-dev \
    && docker-php-ext-configure gd --with-gd --with-webp-dir --with-jpeg-dir \
    --with-png-dir --with-zlib-dir --with-xpm-dir
RUN docker-php-ext-install gd

# PDO
RUN docker-php-ext-install -j$(nproc) pdo

# PDO Mysql
RUN docker-php-ext-install -j$(nproc) pdo_mysql

# Postgres PDO
RUN apt-get update && apt-get install -y libpq-dev \
    && docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install pgsql pdo pdo_pgsql

# Redis
RUN mkdir -p /usr/src/php/ext/redis \
    && curl -L https://github.com/phpredis/phpredis/archive/$PHPREDIS_VERSION.tar.gz | tar xvz -C /usr/src/php/ext/redis --strip 1 \
    && echo 'redis' >> /usr/src/php-available-exts \
    && docker-php-ext-install redis

# INTL
RUN apt-get update && apt-get install -y zlib1g-dev libicu-dev g++ \
  && docker-php-ext-configure intl \
  && docker-php-ext-install intl

# OPCACHE
RUN docker-php-ext-install -j$(nproc) opcache

# APCu
RUN pecl install apcu && docker-php-ext-enable apcu

#composer
RUN apt-get update && \
    apt-get install -y --no-install-recommends git zip unzip

#composer fixes (exhausted memory error) - comment if not sure
RUN touch /usr/local/etc/php/conf.d/memory_limit.ini \
    && echo "memory_limit = 4000M;" >> /usr/local/etc/php/conf.d/memory_limit.ini

#RUN curl --silent --show-error https://getcomposer.org/installer | php
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# xdebug
RUN pecl install xdebug && docker-php-ext-enable xdebug
COPY xdebug.ini /usr/local/etc/php/conf.d/xdebug.ini.tmp
RUN cat /usr/local/etc/php/conf.d/xdebug.ini.tmp >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini
RUN rm /usr/local/etc/php/conf.d/xdebug.ini.tmp

#cron
RUN apt-get install -y cron
ADD crontab /etc/crontab
RUN chmod 0644 /etc/crontab
RUN touch /var/log/cron.log

EXPOSE 9000

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD /entrypoint.sh