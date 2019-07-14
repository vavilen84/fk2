# Symfony 4.* basic application skeleton includes:
- Docker (php + nginx + postgres + adminer + redis)
- Codeception 
- XDebug

## Install Docker 

https://docs.docker.com/install/

## Create the docker group

```
$ sudo groupadd docker
$ sudo usermod -aG docker $USER
```

## Install docker-compose 

https://docs.docker.com/compose/install/

## Install docker-hostmanager

https://github.com/iamluc/docker-hostmanager

run docker-hostmanager
```
$ docker run -d --name docker-hostmanager --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /etc/hosts:/hosts iamluc/docker-hostmanager
```

## Add ENV file

create .env file from .env.dist file and set correct vars values in it

## XDebug

set alias 10.254.254.254 to 127.0.0.1 network interface
```
$ sudo ifconfig lo:0 10.254.254.254 up
```

##  Start with Docker

```
$ docker-compose up -d --build
```

## Install composer libs

```
$ docker exec -it --user 1000 fk2_php_1 composer install
```

## Create database schema

```
$ docker exec -it --user 1000 fk2_php_1 bin/console doctrine:schema:create
```

## Run migrations

```
$ docker exec -it --user 1000 fk2_php_1 bin/console doctrine:migrations:migrate
```

## Load fixtures

```
$ docker exec -it --user 1000 fk2_php_1 bin/console doctrine:fixtures:load
```

## Available URLs:

"http://site.fk2_local/" - website

"http://adminer.fk2_local:8080/" - adminer

Adminer credentials:<br>
System: PostgreSQL<br>
Server: postgres<br>
Username: symfony<br>
Password: 123456

## Codeception

run all tests under folder
```
$ docker exec -it --user 1000 fk2_php_1 bash
$ cd codeception
$ php ../vendor/bin/codecept run tests/Functional
```

run one test in debug mode
```
$ docker exec -it --user 1000 fk2_php_1 bash
$ cd codeception
$ php ../vendor/bin/codecept run tests/Api/HelloFromTestCest.php --debug
```

build tester classes
```
$ docker exec -it --user 1000 fk2_php_1 bash
$ cd codeception
$ php ../vendor/bin/codecept build
```