# Go Gin Prototype

This is a prototype for testing the Gin framework in Go for building a REST API. The prototype simulates a simple hotel reservation system.

## Setup

### Prerequisites

- Install Go from [here](https://golang.org/doc/install)
- Install and setup MySQL from [here](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/)
- Create a `reservation` schema in MySQL
- Create an `api_user` user in MySQL with the password `Test12345!` and grant all privileges on the `reservation` schema

### Install dependencies

```bash
go mod download
```

## Run the app

```bash
go run main.go
```

## Run the tests

To run the tests with code coverage reports to the console, run:

```bash
go test -cover -v ./...
```

### Endpoint testing

Import the `postman_collection.json` file into Postman to test the endpoints. This file can be found in the `tests` directory.

To save the reports to a file, run:

```bash
go test -coverprofile=coverage.out -cover -v ./...
```

To view the coverage report in the browser, run:

```bash
go tool cover -html=coverage.out
```
