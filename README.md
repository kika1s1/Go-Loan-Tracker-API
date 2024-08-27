# Go-Loan-Tracker-API
Go Loan Tracker API
# Loan Tracker API

## Overview

The Loan Tracker API is a backend service built with Golang and the Gin web framework. This API provides functionalities for user management, loan management, authentication, and system logging. It follows clean architecture principles to ensure scalability and maintainability.

## Table of Contents

- [Go-Loan-Tracker-API](#go-loan-tracker-api)
- [Loan Tracker API](#loan-tracker-api)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
  - [Testing](#testing)
  - [API Endpoints](#api-endpoints)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

To set up the project locally, follow these steps:

1. **Clone the repository:**

    ```bash
    git clone https://github.com/kika1s1/Go-Loan-Tracker-API.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd Go-Loan-Tracker-API
    ```

3. **Install the project dependencies:**

    ```bash
    go mod tidy
    ```

4. **Install MongoDB:**

    Follow the [MongoDB installation guide](https://docs.mongodb.com/manual/installation/) for your operating system.

## Configuration

1. **Create a `.env` file in the root directory:**

    ```env
    EMAIL_PROVIDER=example_provider
    SMTP_HOST=smtp.example.com
    EMAIL_PORT=587
    SENDER_EMAIL=example@example.com
    SENDER_PASSWORD=your_password_here
    APP_DOMAIN=localhost:3000
    MONGO_URI=mongodb://127.0.0.1:27017/your_database_name
    JWT_SECRET=your_jwt_secret_here
    GO_PORT=3000
    AI_API_DOMAIN=http://127.0.0.1:8000
    OPENAI_KEY=your_openai_key_here
    GEMINI_API_KEY=your_gemini_api_key_here
    LANGCHAIN_API_KEY=your_langchain_api_key_here
    ACCESSS_TOKEN_LIFE_TIME=3600
    REFRESH_TOKEN_LIFE_TIME=7200
    GOOGLE_CLIENT_ID=your_google_client_id_here
    GOOGLE_CLIENT_SECRET=your_google_client_secret_here
    VERIFICATION_TOKEN_LIFETIME=0.5


    ```

2. **Load environment variables in your `main.go` file:**

    ```go
    import (
        "github.com/joho/godotenv"
        "log"
    )

    func main() {
        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file")
        }
        // Initialize the application
    }
    ```

## Running the Application

1. **Start the MongoDB server:**

    ```bash
    mongod
    ```

2. **Run the Go application:**

    ```bash
    go run main.go
    ```

   By default, the application will start on port 8080. You can adjust the port in the `.env` file if needed.

## Testing

1. **Run unit tests:**

    ```bash
    go test ./...
    ```

2. **End-to-end tests:**

   Use Postman or any HTTP client to test the API endpoints described in the [API Endpoints](#api-endpoints) section.

## API Endpoints

For detailed information on API endpoints, request and response formats, refer to the [API Documentation](https://documenter.getpostman.com/view/36018169/2sAXjGduAB) section in this file.

## Contributing

Contributions are welcome! To contribute to the project:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your forked repository.
5. Open a pull request to the main repository.

Please follow the project's coding style and include tests for new features or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
