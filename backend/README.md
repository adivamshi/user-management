
# User Management API

## Project Structure

- `cmd/server/main.go`: The entry point of the API server.
- `domain/`: Contains the core domain logic.
  - `errors.go`: Defines custom error handling for the domain.
  - `model.go`: Contains the user model definition.
- `repository/`: Data access layer.
  - `database.go`: Setup and utilities for database interaction.
  - `users.go`: Repository methods for user operations.
- `server/server.go`: Server configuration and setup.
- `routes/router.go`: Routing configurations.
- `users/`: Business logic for user operations.
  - `service.go`: Business logic for user services.
  - `endpoint.go`: Defines the endpoints for the user service.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. This project uses Go Modules for dependency management.

### Installing

Install the necessary dependencies:

```bash
go mod tidy
```

### Running the server

To start the server, run:

```bash
go run cmd/server/main.go
```

This will start the server on `http://localhost:8081`. The server endpoints can be accessed from this base URL.

## Usage

Here's how you can interact with the API:

### Create a User

POST `/users`

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}
```

### Get Users

GET `/users`

### Update a User

PUT `/users/{id}`

```json
{
  "name": "John Doe Updated",
  "email": "john.doe@example.com",
  "age": 31
}
```

### Delete a User

DELETE `/users/{id}`