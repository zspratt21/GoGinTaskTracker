# Go Gin DynamoDB API Template

This is a template project for a Golang Gin based REST API that uses DynamoDB as a database.

## Features

- [x] Dockerized local dev environment with Air for hot reloading and a local DynamoDB instance.
- [x] [Gin](https://gin-gonic.com/)
- [x] [AWS SDK](https://github.com/aws/aws-sdk-go)
- [x] Unit tests for JSON Marshalling
- [x] Makefile for common docker tasks and building binaries for different platforms

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Go 1.21.3](https://golang.org/dl/)
- Git Bash or a unix terminal for running commands in the MakeFile

### Installing

Clone the repository or start a new project with this template.

Copy the .env.example file to .env and modify the values as needed.

```bash
cp example.env .env
```

Repeat for example.migrations.env

```bash
cp example.migration.env migration.env
```

### Running the migrations
You can either run the migration script from your IDE or from the make file.
```bash
make run-migrations
```

### Running the tests

```bash
make run-tests
```

### Running the application

You can either launch the docker containers from your IDE or use the start-containers command in the make file.
```bash
make start-containers
```

### Using the example colors crud api

The example api is a simple CRUD api for colors. It has the following endpoints:

- GET /colors - Returns all colors
- GET /colors/:id - Returns a color by id
- POST /colors - Creates a new color
- PUT /colors/:id - Updates a color by id
- DELETE /colors/:id - Deletes a color by id
