# bookstore-users-api

The `bookstore-users-api` is a backend service designed to manage user-related operations in a bookstore application. It provides functionalities for user authentication, registration, profile management, and more, ensuring secure and efficient handling of user data.

## Overview

This project is built using Go (Golang) and leverages several libraries and frameworks to facilitate its operations:

- **Gin-Gonic**: A web framework that helps in building web applications and APIs in Go.
- **Go SQL Driver for MySQL**: Ensures seamless connectivity and interaction with a MySQL database.
- **Godotenv**: Simplifies the management of environment variables by loading them from a `.env` file.

## Project Structure

The project is organized into several directories, each serving a specific purpose:

- **app/**: Contains the main application logic files.
- **data/mysql/**: Contains database-related files.
- **domain/models/**: Contains user-related models.
- **handlers/**: Contains request handlers.
- **services/**: Contains service logic files.
- **utils/**: Contains utility files for handling dates, errors, etc.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Install the required dependencies by running `go mod tidy`.
4. Set up your environment variables in a `.env` file.
5. Run the application using `go run app/app.go`.

## API Endpoints

The API provides several endpoints for managing users. Here are some examples:

- **Create User**: `POST /users` - Registers a new user.
- **Get User**: `GET /users/{id}` - Retrieves a user's details.
- **Update User**: `PUT /users/{id}` - Updates a user's details.
- **Delete User**: `DELETE /users/{id}` - Deletes a user.

## Error Handling

The project uses a custom error handling mechanism to provide more informative error messages. Errors are structured to include a `Message` and `Reason`, making it easier to debug and understand issues.

## Contributing

Contributions to the `bookstore-users-api` are welcome. Please feel free to submit a pull request or open an issue if you encounter any problems.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.