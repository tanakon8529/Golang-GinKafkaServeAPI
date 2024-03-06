# Golang-GinKafkaServeAPI

This project showcases a robust Go-based backend system serving API requests, integrated with Kafka for messaging and Redis for authentication token storage. Leveraging the Gin framework for efficient web routing, the project is neatly documented with Swagger, ensuring ease of use and implementation. Designed for scalability and ease of deployment, the architecture includes multi-container Docker setups for both development and production environments.

## Features

- **Kafka Messaging**: Integrated Kafka producer and consumer routes for efficient message handling.
- **Authentication**: Secure auth routes using tokens stored and managed in Redis.
- **Health Check**: A dedicated health route to monitor the API's status.
- **Swagger Documentation**: All routes are documented with Swagger for easy testing and integration.
- **Middleware Integration**: Middleware for reading Kafka messages and token validation.
- **Docker-compose Support**: Simplified development and deployment with multi-container Docker setups.

## Prerequisites

- Golang
- Gin framework
- Docker
- Docker-compose

## Installation

### Using Docker Compose

1. Clone the repository:
   ```bash
   git clone https://github.com/tanakon8529/Golang-GinKafkaServeAPI.git
   cd Golang-GinKafkaServeAPI
   ```

2. Start the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```

### Local Development

1. Ensure Go is installed and set up on your machine.

2. Install the required Go modules:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## Configuration

Create a `.env` file in the root of your project directory with the following configurations:

Ensure you replace placeholder values with your actual configuration details.

## Running the Project

### Using Docker Compose

To start all services with Docker Compose, run:

```bash
docker-compose up --build
```

### Local Development

For local development, after setting up your `.env` file and installing dependencies, run:

```bash
go run main.go
```

## Usage

To test the API, you can use `curl` or any HTTP client of your choice. Example requests for authentication, health check, and sending messages to Kafka are documented within the project's Swagger UI, accessible at `${HOST}:${PORT_GINAPI}${API_DOC}`.

## Acknowledgments

This project utilizes several Go modules and third-party libraries to provide its functionality:

- Gin for web routing
- Confluent Kafka Go for Kafka integration
- Go-redis for Redis operations
- Swaggo for Swagger documentation
- And many others listed in the `go.mod` file for various functionalities like validation, logging, and environment variable management.

A big thank you to the maintainers of these libraries and to all contributors who have helped in testing, documenting, and enhancing this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details.
