# Golang CRUD Boilerplate

A simple and clean Golang CRUD boilerplate using PostgreSQL as the database.

This project is built using [**Gin**](https://gin-gonic.com/), a fast and lightweight HTTP web framework, and [**Gorm**](https://gorm.io/), an ORM library for Go that simplifies database operations.

## Features

- RESTful API using **Gin**
- PostgreSQL as the database
- Environment configuration using `.env`

## Prerequisites

Ensure you have the following installed:

- **Go**
- **PostgreSQL**

## Installation

1. Clone the repository:
   ```sh
    git clone https://github.com/n-had/golang-crud-boilerplate.git
   ```
   ```sh
   cd golang-crud-boilerplate
   ```
2. Configure the `.env` file

   Check out the .env.example file in the project root directory

3. Install dependencies:
   ```sh
    go mod tidy
   ```
4. Run database migrations
   ```sh
    go run ./migrate/migrate.go
   ```
5. Development Server Setup using [CompileDaemon](https://github.com/githubnemo/CompileDaemon):

   CompileDaemon watches your Go files and automatically recompiles and restarts your server when files change.

   ```sh
   go install github.com/githubnemo/CompileDaemon@latest
   ```

   Start the Development Server

   ```sh
   CompileDaemon -command="./golang-crud-boilerplate"

   ```

6. Test Endpoints:
   Use an API client (e.g., Postman or Insomnia) to test the endpoints of the running server.
