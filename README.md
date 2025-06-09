# Oolio Backend Challenge

This project is a Go-based REST API for a food ordering system. It provides endpoints for managing products, orders, and promo code validation.

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 12 or higher
- Make

## Setup

1. **Clone the repository:**
   ```sh
   git clone <repository-url>
   cd oolio-backend-challenge
   ```

2. **Set up the database:**
   - Create a PostgreSQL database named `food_ordering`.
   - Run the migrations to create the necessary tables and seed data:
     ```sh
     make migrate
     ```

3. **Build the project:**
   ```sh
   make build
   ```

4. **Run the server:**
   ```sh
   PORT=9090 make run
   ```
   The server will start on port 9090 (or the port specified by the `PORT` environment variable).

## API Endpoints

### Products

- **GET /product**
  - Returns a list of all products.
  - Requires a valid API key in the `api_key` header.

- **GET /product/:id**
  - Returns a specific product by ID.
  - Requires a valid API key in the `api_key` header.

### Orders

- **GET /orders**
  - Returns a list of all orders.
  - Requires a valid API key in the `api_key` header.

- **GET /order/:id**
  - Returns a specific order by ID.
  - Requires a valid API key in the `api_key` header.

- **POST /order**
  - Creates a new order.
  - Requires a valid API key in the `api_key` header.
  - Request body example:
    ```json
    {
      "items": [
        { "productId": "1", "quantity": 2 },
        { "productId": "3", "quantity": 1 }
      ]
    }
    ```

### Promo Codes

- **GET /validate-promo/:code**
  - Validates a promo code.
  - Requires a valid API key in the `api_key` header.
  - Returns `{"valid": true}` if the code is valid, `{"valid": false}` otherwise.

## Testing

Run the test script to verify all API endpoints:
```sh
./scripts/test_api.sh
```

## Project Structure

- **cmd/api/main.go**: Entry point for the API server.
- **internal/domain**: Domain models and interfaces.
- **internal/handler**: HTTP handlers for API endpoints.
- **internal/repository**: Database repositories (PostgreSQL).
- **internal/service**: Business logic for products, orders, and promo codes.
- **migrations**: SQL migrations for database setup and seeding.

## License

This project is licensed under the MIT License. See the LICENSE file for details.