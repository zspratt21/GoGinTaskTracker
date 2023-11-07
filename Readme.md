# Go Gin DynamoDB API Template
An API for tracking tasks, built using Golang and the Gin web framework.

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

Use an API client such as ARC or Postman, to make requests to the API. By default, the API is hosted at http://localhost:9001. The following endpoints are available:
#### GET /api/Tasks
Returns a list of all tasks.
#### GET /api/Tasks/{id}
Returns a single task with the specified id.
#### POST /api/Tasks
Creates a new task. The request body should contain a JSON object with the following properties:
- name (string)
- isCompleted (bool)
#### PUT /api/Tasks/{id}
Updates an existing task with the specified id. The request body should contain a JSON object with the following properties:
- id (int)
- name (string)
- isCompleted (bool)
#### DELETE /api/Tasks/{id}
Deletes an existing task with the specified id.

## Made With
- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)
- [DynamoDB](https://aws.amazon.com/dynamodb/)
- [Docker](https://www.docker.com/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Goland](https://www.jetbrains.com/go/)