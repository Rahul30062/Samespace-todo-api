# Samespace-todo-api
# TODO API with Golang and ScyllaDB

This project is a TODO API built using Golang and ScyllaDB, containerized and managed using Docker.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Prerequisites
- [Docker](https://www.docker.com/get-started) and Docker Compose
- [Golang](https://golang.org/dl/)
- A ScyllaDB instance (can be set up using Docker)

## Installation
#Step 1: Set Up ScyllaDB with Docker
version: '3.7'

services:
  scylla:
    image: scylladb/scylla
    ports:
      - "9042:9042"
    volumes:
      - scylla-data:/var/lib/scylla

#Step 2: Configure the Database
docker exec -it scylla-container cqlsh

CREATE KEYSPACE todo_api WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
USE todo_api;

CREATE TABLE todos (
    id UUID PRIMARY KEY,
    user_id UUID,
    title TEXT,
    description TEXT,
    status TEXT,
    created TIMESTAMP,
    updated TIMESTAMP
);
#Step 3: Run Go Application
go mod tidy
go run main.go


## To Run application, go to this web point-
# For Creating todos- 127.0.0.1:8080/todos
# For Get all Todos- 127.0.0.1:8080/todos
# For Updating Todos- 127.0.0.1:8080/todos/{id}
# For Deleting Todos- 127.0.0.1:8080/todos/{id}
