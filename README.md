# FizzBuzz REST API Server

A production-ready REST API server that implements a customizable FizzBuzz algorithm with request statistics tracking.

## Features

- **Customizable FizzBuzz**: Configure your own multiples and replacement strings
- **Statistics Tracking**: Track the most frequently requested parameters
- **Production Ready**: Includes Docker support, health checks, and proper error handling

## API Endpoints

### GET /api/v1/fizzbuzz

Generates a FizzBuzz sequence based on the provided parameters.

**Query Parameters:**
- `int1` (required): First integer for replacement logic
- `int2` (required): Second integer for replacement logic  
- `limit` (required): Upper limit for the sequence (1 to 10000)
- `str1` (required): String to replace multiples of int1
- `str2` (required): String to replace multiples of int2

**Example Request:**
```
GET /api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz
```

**Example Response:**
```json
{
  "result": ["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
}
```

### GET /api/v1/stats

Returns statistics about the most frequently requested parameters.

**Example Response:**
```json
{
  "int1": 3,
  "int2": 5,
  "limit": 100,
  "str1": "fizz",
  "str2": "buzz",
  "count": 42
}
```

### GET /health

Health check endpoint for monitoring.

**Example Response:**
```json
{
  "status": "healthy"
}
```

## Algorithm Logic

For each number from 1 to `limit`:
- If the number is divisible by both `int1` and `int2`: replace with `str1str2`
- If the number is divisible by `int1`: replace with `str1`
- If the number is divisible by `int2`: replace with `str2`
- Otherwise: keep the number as is

## Running the Application

### Local Development

1. **Prerequisites:**
   - Go 1.24.5 or later

2. **Clone and run:**
   ```bash
   go mod tidy
   go run main.go
   ```

3. **The server will start on port 8080 by default**

### Using Docker

1. **Build the image:**
   ```bash
   docker build -t fizzbuzz-api .
   ```

2. **Run the container:**
   ```bash
   docker run -p 8080:8080 fizzbuzz-api
   ```

### Environment Variables

- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin framework mode (default: release)

## Testing

Run the test suite:
```bash
go test ./...
```

## API Documentation

The project uses Swagger/OpenAPI 3.0 for API documentation. To regenerate the documentation:

```bash
go run github.com/swaggo/swag/cmd/swag@latest init
```

Access the interactive documentation at `http://localhost:8080/swagger/index.html` when running the server.

## Architecture

The application follows a clean architecture pattern:

- **Handler Layer**: Handles HTTP requests/responses, input validation, and statistics retrieval
- **Service Layer**: Contains business logic for FizzBuzz generation and request statistics

### Statistics Storage

The request statistics are currently stored in memory for simplification. In a production environment, these statistics should be persisted in a database (e.g., PostgreSQL, MongoDB, or Redis) to ensure data durability across server restarts and to support horizontal scaling.

## Production Considerations

- **Thread Safety**: Statistics are tracked using mutex locks
- **Input Validation**: All parameters are validated with appropriate limits
- **Error Handling**: Comprehensive error responses with meaningful messages
- **Health Checks**: Built-in health endpoint for monitoring
- **Security**: Non-root user in Docker container
- **Resource Limits**: Request limit capped at 10,000 to prevent abuse

## Example Usage Locally

```bash
# View interactive API documentation
open http://localhost:8080/swagger/index.html

# Basic fizzbuzz
curl "http://localhost:8080/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"

# Custom parameters
curl "http://localhost:8080/api/v1/fizzbuzz?int1=2&int2=7&limit=20&str1=foo&str2=bar"

# Check statistics
curl "http://localhost:8080/api/v1/stats"

# Health check
curl "http://localhost:8080/health"
```