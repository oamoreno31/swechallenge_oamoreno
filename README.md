# swechallenge_oamoreno
Custom platform for swechallenge.

## Backend Overview

The backend is a RESTful API built with Go (Golang) using the Gin web framework. It provides a complete CRUD (Create, Read, Update, Delete) interface for managing financial items/products, including stock tickers, company information, brokerage ratings, and target prices.

### Key Features
- RESTful API endpoints for product management
- PostgreSQL database integration
- JSON request/response handling
- Input validation using Gin binding
- Automatic database table creation
- Environment-based configuration

## Project Structure

```
backend/
├── main.go                 # Application entry point, server initialization
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── config/
│   └── database.go        # Database connection and table creation
├── controllers/
│   └── item_controller.go # Business logic and HTTP handlers
├── models/
│   └── item.go            # Data models and structs
├── routes/
│   └── routes.go          # API route definitions
└── learning/              # Learning/example files
    ├── first.go
    └── main.go
```

### Directory Responsibilities

- **config/**: Database connection management and schema initialization
- **controllers/**: HTTP request handlers implementing business logic
- **models/**: Data structures and validation rules
- **routes/**: API endpoint definitions and routing configuration
- **main.go**: Application bootstrap, environment loading, and server startup

## Functionality

### API Endpoints

All endpoints are prefixed with `/api/products`:

| Method | Endpoint | Description | Status Codes |
|--------|----------|-------------|--------------|
| `GET` | `/api/products` | Retrieve all products | 200, 500 |
| `GET` | `/api/products/:id` | Get a specific product by ID | 200, 404, 500 |
| `POST` | `/api/products` | Create a new product | 201, 400, 500 |
| `PUT` | `/api/products/:id` | Update an existing product | 200, 400, 404, 500 |
| `DELETE` | `/api/products/:id` | Delete a product | 200, 404, 500 |

### Data Model

The `Item` model represents a financial product with the following fields:

- `ID` (int64): Unique identifier (auto-generated)
- `Ticker` (string, required): Stock ticker symbol
- `Target_from` (string, required): Lower target price
- `Target_to` (string, required): Upper target price
- `Company` (string, required): Company name
- `Action` (string, required): Recommended action
- `Brokerage` (string, required): Brokerage firm name
- `Rating_from` (string, required): Previous rating
- `Rating_to` (string, required): New rating
- `Time` (string, required): Timestamp or time information

## Installation

### Prerequisites

- Go 1.25.5 or higher
- PostgreSQL database
- Git

### Setup Steps

1. **Clone the repository** (if not already done):
   ```bash
   git clone <repository-url>
   cd swechallenge_oamoreno/backend
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Configure environment variables**:
   Create a `.env` file in the `backend/` directory with the following variables:
   ```env
   DATABASE_URL_CONECTION=postgres://username:password@localhost:5432/dbname?sslmode=disable
   PORT=8080
   ```
   
   Alternatively, set these as system environment variables.

4. **Start PostgreSQL**:
   Ensure your PostgreSQL server is running and accessible.

5. **Run the application**:
   ```bash
   go run main.go
   ```

   The server will start on port 8080 (or the port specified in the `PORT` environment variable).

6. **Verify the connection**:
   The application will automatically:
   - Connect to the database
   - Create the `items` table if it doesn't exist
   - Display the current database timestamp on successful connection

## Usage

### Starting the Server

```bash
cd backend
go run main.go
```

The server will be available at `http://localhost:8080`

### API Usage Examples

#### 1. Get All Products
```bash
curl http://localhost:8080/api/products
```

#### 2. Get a Specific Product
```bash
curl http://localhost:8080/api/products/1
```

#### 3. Create a New Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "Ticker": "AAPL",
    "Target_from": "150",
    "Target_to": "180",
    "Company": "Apple Inc.",
    "Action": "Buy",
    "Brokerage": "Goldman Sachs",
    "Rating_from": "Hold",
    "Rating_to": "Buy",
    "Time": "2024-01-15T10:30:00Z"
  }'
```

#### 4. Update a Product
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "Ticker": "AAPL",
    "Target_from": "160",
    "Target_to": "190",
    "Company": "Apple Inc.",
    "Action": "Strong Buy",
    "Brokerage": "Goldman Sachs",
    "Rating_from": "Buy",
    "Rating_to": "Strong Buy",
    "Time": "2024-01-16T10:30:00Z"
  }'
```

#### 5. Delete a Product
```bash
curl -X DELETE http://localhost:8080/api/products/1
```

### Response Format

**Success Response (200/201)**:
```json
{
  "ID": 1,
  "Ticker": "AAPL",
  "Target_from": "150",
  "Target_to": "180",
  "Company": "Apple Inc.",
  "Action": "Buy",
  "Brokerage": "Goldman Sachs",
  "Rating_from": "Hold",
  "Rating_to": "Buy",
  "Time": "2024-01-15T10:30:00Z"
}
```

**Error Response (400/404/500)**:
```json
{
  "error": "Error message description"
}
```

## User Testing Plan

### Test Scenarios

#### 1. **Product Creation Tests**
- ✅ **TC-001**: Create a product with all required fields
  - Expected: Status 201, product returned with generated ID
- ✅ **TC-002**: Create a product with missing required fields
  - Expected: Status 400, error message indicating missing fields
- ✅ **TC-003**: Create a product with invalid JSON
  - Expected: Status 400, JSON parsing error

#### 2. **Product Retrieval Tests**
- ✅ **TC-004**: Get all products (empty database)
  - Expected: Status 200, empty array `[]`
- ✅ **TC-005**: Get all products (with data)
  - Expected: Status 200, array of products
- ✅ **TC-006**: Get product by valid ID
  - Expected: Status 200, product object
- ✅ **TC-007**: Get product by invalid ID (non-existent)
  - Expected: Status 404, "Producto no encontrado"

#### 3. **Product Update Tests**
- ✅ **TC-008**: Update product with valid ID and complete data
  - Expected: Status 200, updated product
- ✅ **TC-009**: Update product with invalid ID
  - Expected: Status 404, "Producto no encontrado"
- ✅ **TC-010**: Update product with missing required fields
  - Expected: Status 400, validation error

#### 4. **Product Deletion Tests**
- ✅ **TC-011**: Delete product with valid ID
  - Expected: Status 200, success message
- ✅ **TC-012**: Delete product with invalid ID
  - Expected: Status 404, "Producto no encontrado"

#### 5. **Integration Tests**
- ✅ **TC-013**: Complete CRUD workflow
  1. Create product → Verify creation
  2. Get product by ID → Verify data
  3. Update product → Verify changes
  4. Get all products → Verify product appears
  5. Delete product → Verify deletion
  6. Get product by ID → Verify 404

#### 6. **Edge Cases**
- ✅ **TC-014**: Very long string values
  - Test with maximum length strings for each field
- ✅ **TC-015**: Special characters in fields
  - Test with special characters, unicode, etc.
- ✅ **TC-016**: Concurrent requests
  - Test multiple simultaneous create/update operations

### Testing Tools

Recommended tools for testing:
- **Postman**: For manual API testing and collection management
- **curl**: Command-line testing (examples provided above)
- **HTTPie**: User-friendly CLI HTTP client
- **Automated Testing**: Go testing framework for unit/integration tests

### Test Data Examples

```json
{
  "Ticker": "MSFT",
  "Target_from": "300",
  "Target_to": "350",
  "Company": "Microsoft Corporation",
  "Action": "Buy",
  "Brokerage": "Morgan Stanley",
  "Rating_from": "Hold",
  "Rating_to": "Buy",
  "Time": "2024-01-20T14:00:00Z"
}
```

### Expected Test Results

All test cases should pass with:
- Correct HTTP status codes
- Proper JSON response format
- Accurate error messages
- Data persistence in database
- No server crashes or panics

---

## Development Notes

- The application uses Gin framework for HTTP routing and JSON handling
- Database connection uses pgx driver for PostgreSQL
- Environment variables are loaded using godotenv
- The database table is automatically created on first run
- All endpoints return JSON responses
- Error handling includes proper HTTP status codes and error messages