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

## Frontend Overview

The frontend is a modern single-page application (SPA) built with Vue 3, TypeScript, Pinia, and TailwindCSS. It provides an intuitive user interface for managing financial products with a dark blue theme and responsive design.

### Key Features
- 🎨 Dark blue themed interface with gradient backgrounds
- 📊 Interactive dashboard with statistics and recent products
- 📋 Complete product list with full CRUD operations
- ➕ Create product form with validation
- ✏️ Update product form with pre-filled data
- 🗑️ Delete functionality with confirmation dialogs
- 🔄 Real-time state management with Pinia
- 🎯 TypeScript for type safety
- 📱 Fully responsive design
- ⚡ Fast and optimized with Vite

## Frontend Project Structure

```
frontend/sweui/
├── index.html              # HTML entry point
├── package.json            # Node.js dependencies
├── vite.config.ts          # Vite configuration
├── tailwind.config.js      # TailwindCSS theme configuration
├── postcss.config.js       # PostCSS configuration
├── tsconfig.json           # TypeScript configuration
├── src/
│   ├── main.ts             # Application entry point
│   ├── App.vue             # Root component with navigation
│   ├── assets/
│   │   └── main.css        # Global styles and Tailwind imports
│   ├── router/
│   │   └── index.ts        # Vue Router configuration
│   ├── services/
│   │   └── api.ts          # API service for backend communication
│   ├── stores/
│   │   └── productStore.ts # Pinia store for state management
│   ├── types/
│   │   └── product.ts      # TypeScript interfaces
│   └── views/
│       ├── Dashboard.vue  # Dashboard with statistics
│       ├── ProductList.vue # Product listing with delete
│       ├── CreateProduct.vue # Create product form
│       └── UpdateProduct.vue # Update product form
└── public/                 # Static assets
```

### Directory Responsibilities

- **router/**: Client-side routing configuration with Vue Router
- **services/**: API communication layer using Axios
- **stores/**: Global state management with Pinia
- **types/**: TypeScript type definitions and interfaces
- **views/**: Page-level Vue components
- **assets/**: Global CSS and styling resources

## Frontend Functionality

### Routes

| Route | Component | Description |
|-------|-----------|-------------|
| `/` | Dashboard | Main dashboard with statistics and recent products |
| `/products` | ProductList | Full list of all products with delete functionality |
| `/products/create` | CreateProduct | Form to create a new product |
| `/products/edit/:id` | UpdateProduct | Form to edit an existing product |

### Components Overview

#### Dashboard (`/`)
- Displays total product count
- Shows recent products (last 5)
- Quick action buttons
- System status indicator
- Responsive card-based layout

#### Product List (`/products`)
- Complete table view of all products
- Sortable columns
- Delete button with confirmation
- Edit button for each product
- Empty state handling
- Loading and error states

#### Create Product (`/products/create`)
- Form with all required fields
- Real-time validation
- Error message display
- Success redirect to product list
- Cancel button to return to list

#### Update Product (`/products/edit/:id`)
- Pre-filled form with existing data
- Same validation as create form
- Updates existing product
- Error handling for invalid IDs

### State Management

The application uses Pinia for state management with the following store:

**`productStore`**:
- `products`: Array of all products
- `selectedProduct`: Currently selected product
- `loading`: Loading state indicator
- `error`: Error message state
- `productCount`: Computed total count
- Actions: `fetchProducts`, `fetchProductById`, `createProduct`, `updateProduct`, `deleteProduct`

## Frontend Installation

### Prerequisites

- Node.js 20.19.0 or >=22.12.0
- npm or yarn package manager
- Backend API running on `http://localhost:8080`

### Setup Steps

1. **Navigate to frontend directory**:
   ```bash
   cd frontend/sweui
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Configure environment variables**:
   Create a `.env` file in the `frontend/sweui/` directory:
   ```env
   VITE_API_BASE_URL=http://localhost:8080/api
   ```
   
   If the backend runs on a different port or URL, update this accordingly.

4. **Start the development server**:
   ```bash
   npm run dev
   ```

   The application will be available at `http://localhost:5173` (or the port shown in the terminal).

5. **Build for production** (optional):
   ```bash
   npm run build
   ```

   The production build will be in the `dist/` directory.

## Frontend Usage

### Development Mode

```bash
cd frontend/sweui
npm run dev
```

The development server will:
- Start on `http://localhost:5173` (or next available port)
- Enable hot module replacement (HMR)
- Show compilation errors in the browser
- Provide source maps for debugging

### Available Scripts

- `npm run dev` - Start development server with HMR
- `npm run build` - Build for production
- `npm run preview` - Preview production build locally
- `npm run type-check` - Run TypeScript type checking

### User Interface Features

#### Navigation
- Top navigation bar with dark blue theme
- Active route highlighting
- Responsive mobile menu

#### Color Scheme
- **Primary Color**: `#0066e6` (dark blue)
- **Background**: Gradient from `primary-900` to `primary-700`
- **Cards**: Semi-transparent white with backdrop blur
- **Text**: White and light blue shades for contrast

#### Form Validation
- Required field indicators (red asterisk)
- Real-time validation feedback
- Error messages displayed in red
- Disabled submit button during loading

#### Loading States
- Spinner animations during API calls
- Loading text indicators
- Disabled buttons to prevent duplicate submissions

#### Error Handling
- Error messages displayed in red cards
- Retry buttons for failed operations
- User-friendly error messages
- Console logging for debugging

### API Integration

The frontend communicates with the backend through:

- **Base URL**: Configurable via `VITE_API_BASE_URL` environment variable
- **Default**: `http://localhost:8080/api`
- **Endpoints**: All CRUD operations on `/products`

#### API Service Methods

```typescript
// Get all products
productService.getAllProducts(): Promise<Product[]>

// Get product by ID
productService.getProductById(id: number): Promise<Product>

// Create new product
productService.createProduct(data: ProductFormData): Promise<Product>

// Update product
productService.updateProduct(id: number, data: ProductFormData): Promise<Product>

// Delete product
productService.deleteProduct(id: number): Promise<void>
```

### Styling

The application uses **TailwindCSS** with a custom dark blue theme:

- **Primary Colors**: `primary-50` through `primary-900`
- **Main Color**: `primary-500` (#0066e6)
- **Utilities**: Full Tailwind utility classes
- **Custom Styles**: Global styles in `src/assets/main.css`
- **Responsive**: Mobile-first responsive design

### Technologies Used

- **Vue 3.5.25**: Progressive JavaScript framework with Composition API
- **TypeScript 5.9**: Type-safe JavaScript
- **Pinia 3.0.4**: State management library
- **Vue Router 4.4.5**: Client-side routing
- **TailwindCSS 3.4.17**: Utility-first CSS framework
- **Axios 1.7.9**: HTTP client for API requests
- **Vite 7.2.4**: Build tool and development server

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

## Quick Start Guide

### Running the Full Stack Application

1. **Start the Backend**:
   ```bash
   cd backend
   go run main.go
   ```
   Backend will run on `http://localhost:8080`

2. **Start the Frontend** (in a new terminal):
   ```bash
   cd frontend/sweui
   npm install  # First time only
   npm run dev
   ```
   Frontend will run on `http://localhost:5173`

3. **Access the Application**:
   - Open your browser and navigate to `http://localhost:5173`
   - The frontend will automatically connect to the backend API

### Environment Setup

**Backend** (`.env` in `backend/`):
```env
DATABASE_URL_CONECTION=postgres://username:password@localhost:5432/dbname?sslmode=disable
PORT=8080
```

**Frontend** (`.env` in `frontend/sweui/`):
```env
VITE_API_BASE_URL=http://localhost:8080/api
```

---

## Development Notes

### Backend
- The application uses Gin framework for HTTP routing and JSON handling
- Database connection uses pgx driver for PostgreSQL
- Environment variables are loaded using godotenv
- The database table is automatically created on first run
- Test data is automatically seeded if the table is empty
- CORS is configured to allow frontend connections
- All endpoints return JSON responses
- Error handling includes proper HTTP status codes and error messages
- ID fields are converted from database integers to strings for API responses

### Frontend
- Built with Vue 3 Composition API and TypeScript
- State management handled by Pinia stores
- Routing managed by Vue Router
- API communication via Axios with error handling
- Styling with TailwindCSS and custom dark blue theme
- Responsive design for mobile and desktop
- Form validation and user feedback
- Loading states and error handling throughout
- Hot Module Replacement (HMR) enabled in development
- Type-safe development with TypeScript