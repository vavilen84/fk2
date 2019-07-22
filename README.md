# Golang (Beego framework) test application

Application includes:
- basic JWT auth implementation
- CRUD (create/update/edit/delete post)
- Models Validation
- Docker (mysql, phpmyadmin, application containers)
- Tests
- Migrations
- Console command shortcuts
- Modules


## Install Docker 

https://docs.docker.com/install/

## Install docker-compose 

https://docs.docker.com/compose/install/

## Install docker-hostmanager

https://github.com/iamluc/docker-hostmanager

Run manager

```
$ docker run -d --name docker-hostmanager --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /etc/hosts:/hosts iamluc/docker-hostmanager
```

Beego config don`t support .env vars, so in current implementation all vars (paths, db names, etc) hardcoded. 
So you need to clone project in /var/www/go-test-app folder to work correct. May be fixed in future.

## Start with Docker

```
$ cd /project/path
$ docker compose up -d --build
```

## Add modules

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# go mod tidy
```

or

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# bee rs load:modules
```

## Create application database schema

```
cd /project/path
docker exec -it fk2_baseapp_1 bash
bee migrate -conn="root:123456@tcp(db:3306)/godb"
```

or

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# bee rs migrate:db
```

## Create database schema for tests

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# bee migrate -conn="root:123456@tcp(db:3306)/godb_test"
```

or

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# bee rs migrate:test_db
```

## Run tests

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# go test ./...
```

or 

```
$ cd /project/path
$ docker exec -it go-test-app_baseapp_1 bash
# bee rs test
```

## Available URLs
 
Application: http://baseapp.go-test-app_local 

PHPMyAdmin: http://pma.go-test-app_local





