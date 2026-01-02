<font size= "5"> **Table Of Contents** </font>
- [Todo list](#todo-list)
- [Introduction](#introduction)
- [Environment for development](#environment-for-development)
  - [Mysql](#mysql)
  - [Service](#service)
    - [Preparation](#preparation)
    - [Run service as standalone](#run-service-as-standalone)
    - [Database migration](#database-migration)
- [Testing](#testing)
- [Contributors](#contributors)


# Todo list
- [x] Integrate mock framework for automation testing

# Introduction
- A microservice for managing accounts belongs to Fiagram project.

# Environment for development
## Mysql
- Start MySQL database by docker with the default configuration from `configs/local.yaml`
```
./scripts/run-docker-mysql-dev.sh
```
## Service
### Preparation
- Download tools for the service
```
make init
```
### Run service as standalone
- Run service with the make command
```
make run-standalone-server
```
### Database migration
- Add a new schema for database
```
make migrate-new <name_of_new_schema>
```
# Testing
- Require to the Mysql database must be run at first [Run MySQL with docker](#mysql)
```
make test
```

# Contributors
|  Fullname   | Startdate  |
| :---------: | :--------: |
| Vũ Đức Thái | 01/01/2026 |


