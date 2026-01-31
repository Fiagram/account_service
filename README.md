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
- [Deployment](#deployment)


# Todo list
- [x] Integrate mock framework for automation testing
- [ ] Validate input to grpc server by extended protobuf package
- [x] Fixbug SQL injection when querying database directly
- [x] Apply Docker for production releases

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

# Deployment
- Build a docker image with the following command:
```
./scripts/build-docker-image-deployment.sh
```
- After building the docker image for `account_service`, let's try to deploy services (including mysql database):
```
make docker-compose-up-test
```
