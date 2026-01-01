<font size= "5"> **Table Of Contents** </font>
- [Todo list](#todo-list)
- [Introduction](#introduction)
- [Environment for development](#environment-for-development)
  - [Mysql](#mysql)
  - [Service](#service)
    - [Database migration](#database-migration)
- [Contributors](#contributors)


# Todo list
- [ ] Integrate mock framework for automation testing

# Introduction
- A microservice for managing accounts belongs to Fiagram project.

# Environment for development
## Mysql
- Start MySQL database by docker
```
./scripts/run-docker-mysql-dev.sh
```
## Service
- Download tools for service
```
make init
```
- Run service with the make command
```
make run-standalone-server
```
### Database migration
- Add a new schema for database
```
make migrate-new <name_of_new_schema>
```


# Contributors
|  Fullname   | Startdate  |
| :---------: | :--------: |
| Vũ Đức Thái | 01/01/2026 |


