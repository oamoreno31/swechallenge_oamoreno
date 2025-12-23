# Frontend Application

Vue 3 + TypeScript + Pinia + TailwindCSS application for managing financial products.

## Features

- 🎨 Dark blue themed interface
- 📊 Dashboard with statistics and recent products
- 📋 Product list with full CRUD operations
- ➕ Create product form
- ✏️ Update product form
- 🗑️ Delete functionality with confirmation
- 🔄 Real-time state management with Pinia
- 🎯 TypeScript for type safety

## Project Structure

```
src/
├── assets/
│   └── main.css          # TailwindCSS imports and global styles
├── router/
│   └── index.ts          # Vue Router configuration
├── services/
│   └── api.ts            # API service for backend communication
├── stores/
│   └── productStore.ts   # Pinia store for product state management
├── types/
│   └── product.ts        # TypeScript interfaces
├── views/
│   ├── Dashboard.vue     # Main dashboard view
│   ├── ProductList.vue    # Product listing with delete
│   ├── CreateProduct.vue  # Create product form
│   └── UpdateProduct.vue  # Update product form
├── App.vue                # Root component with navigation
└── main.ts                # Application entry point
```

## Setup

### Prerequisites

- Node.js 20.19.0 or >=22.12.0
- npm or yarn

### Installation

1. Install dependencies:
```bash
npm install
```

2. Create a `.env` file in the root directory:
```env
VITE_API_BASE_URL=http://localhost:8080/api
```

3. Start the development server:
```bash
npm run dev
```

The application will be available at `http://localhost:5173` (or the port shown in the terminal).

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run type-check` - Run TypeScript type checking

## Routes

- `/` - Dashboard
- `/products` - Product list
- `/products/create` - Create new product
- `/products/edit/:id` - Edit existing product

## API Configuration

The application expects the backend API to be running on `http://localhost:8080/api` by default. You can change this by setting the `VITE_API_BASE_URL` environment variable.

## Styling

The application uses TailwindCSS with a custom dark blue color scheme:
- Primary color: `#0066e6` (primary-500)
- Dark backgrounds: primary-900, primary-800, primary-700
- Light accents: primary-200, primary-300

## Technologies

- **Vue 3** - Progressive JavaScript framework
- **TypeScript** - Type-safe JavaScript
- **Pinia** - State management
- **Vue Router** - Client-side routing
- **TailwindCSS** - Utility-first CSS framework
- **Axios** - HTTP client
- **Vite** - Build tool and dev server
