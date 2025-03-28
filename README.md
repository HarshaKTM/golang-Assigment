# Book Management System

A complete book management application with a Go-based RESTful API backend and a React-based frontend. This application provides CRUD operations for books, an optimized search endpoint, and a modern user interface.

## Project Structure

- **Backend**: Go-based RESTful API with MongoDB and file-based storage options
- **Frontend**: React application with Material UI components

## Features

### Backend

- CRUD operations for books
- MongoDB and file-based storage options
- Optimized search using Go's concurrency features (goroutines and channels)
- Dockerized application
- Kubernetes deployment support

### Frontend

- Modern UI built with Material-UI components
- Complete CRUD operations for books
- Real-time search with optimized backend integration
- Responsive design works on all device sizes
- Form validation for data integrity
- Notification system for user feedback

## Requirements

- Go 1.20 or higher
- Node.js 14 or higher
- npm or yarn
- MongoDB (optional)
- Docker (optional)

## Step-by-Step Installation

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/book-management-system.git
cd book-management-system
```

### 2. Install Backend Dependencies

```bash
go mod download
go mod tidy
```

### 3. Install Frontend Dependencies

```bash
cd frontend
npm install
cd ..
```

## Running the Application

### Quick Start (Both Frontend and Backend)

Use the provided batch script to start both services:
```bash
start_app.bat
```

This will:
1. Start the backend API on port 5001 with sample data
2. Start the frontend application on port 3000

### Backend Development

1. Run the application with file storage:
   ```bash
   go run main.go -port 5001
   ```
2. Run with MongoDB:
   ```bash
   go run main.go -port 5001 -mongodb
   ```
3. Run with sample data:
   ```bash
   go run main.go -port 5001 -seed
   ```

### Frontend Development

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Start the development server:
   ```bash
   npm start
   ```

## Testing API Endpoints

Below are examples of how to test each CRUD endpoint using curl commands:

### 1. List All Books (GET /books)

```bash
curl http://localhost:5001/books
```

Expected Response:
```json
[
  {
    "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
    "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
    "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
    "title": "The Great Gatsby",
    "publicationDate": "1925-04-10",
    "isbn": "9780743273565",
    "pages": 180,
    "genre": "Novel",
    "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
    "price": 15.99,
    "quantity": 5
  },
  // More books...
]
```

### 2. Create a New Book (POST /books)

```bash
curl -X POST http://localhost:5001/books \
  -H "Content-Type: application/json" \
  -d '{
    "authorId": "author-123",
    "publisherId": "publisher-456",
    "title": "New Book Title",
    "publicationDate": "2023-01-15",
    "isbn": "9781234567890",
    "pages": 250,
    "genre": "Science Fiction",
    "description": "A brand new science fiction novel about future technologies.",
    "price": 19.99,
    "quantity": 10
  }'
```

Expected Response:
```json
{
  "bookId": "generated-id-here",
  "authorId": "author-123",
  "publisherId": "publisher-456",
  "title": "New Book Title",
  "publicationDate": "2023-01-15",
  "isbn": "9781234567890",
  "pages": 250,
  "genre": "Science Fiction",
  "description": "A brand new science fiction novel about future technologies.",
  "price": 19.99,
  "quantity": 10
}
```

### 3. Get a Book by ID (GET /books/{id})

First, get a book ID from the list or create a new book. Then:

```bash
curl http://localhost:5001/books/bb329a31-6b1e-4daa-87ee-71631aa05866
```

Expected Response:
```json
{
  "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
  "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
  "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
  "title": "The Great Gatsby",
  "publicationDate": "1925-04-10",
  "isbn": "9780743273565",
  "pages": 180,
  "genre": "Novel",
  "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
  "price": 15.99,
  "quantity": 5
}
```

### 4. Update a Book (PUT /books/{id})

```bash
curl -X PUT http://localhost:5001/books/bb329a31-6b1e-4daa-87ee-71631aa05866 \
  -H "Content-Type: application/json" \
  -d '{
    "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
    "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
    "title": "The Great Gatsby (Updated)",
    "publicationDate": "1925-04-10",
    "isbn": "9780743273565",
    "pages": 180,
    "genre": "Classic Literature",
    "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
    "price": 17.99,
    "quantity": 8
  }'
```

Expected Response:
```json
{
  "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
  "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
  "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
  "title": "The Great Gatsby (Updated)",
  "publicationDate": "1925-04-10",
  "isbn": "9780743273565",
  "pages": 180,
  "genre": "Classic Literature",
  "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
  "price": 17.99,
  "quantity": 8
}
```

### 5. Delete a Book (DELETE /books/{id})

```bash
curl -X DELETE http://localhost:5001/books/bb329a31-6b1e-4daa-87ee-71631aa05866
```

Expected Response: HTTP 204 No Content (no body)

### 6. Search Books (GET /books/search?q=<keyword>)

```bash
curl "http://localhost:5001/books/search?q=gatsby"
```

Expected Response:
```json
[
  {
    "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
    "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
    "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
    "title": "The Great Gatsby",
    "publicationDate": "1925-04-10",
    "isbn": "9780743273565",
    "pages": 180,
    "genre": "Novel",
    "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
    "price": 15.99,
    "quantity": 5
  }
]
```

## Testing with PowerShell Script

For Windows users, you can run the included PowerShell script to test all endpoints:

```powershell
.\test_api.ps1
```

## Testing with Postman

1. Download and install [Postman](https://www.postman.com/downloads/)
2. Import the collection file: `Book_API.postman_collection.json` (if available)
3. Or create a new collection with the endpoints described above

## Book Model

```json
{
  "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
  "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
  "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
  "title": "The Great Gatsby",
  "publicationDate": "1925-04-10",
  "isbn": "9780743273565",
  "pages": 180,
  "genre": "Novel",
  "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
  "price": 15.99,
  "quantity": 5
}
```

## Docker Deployment Steps

### Backend Docker Deployment

1. Build the backend Docker image:
   ```bash
   docker build -t book-api .
   ```
2. Run the container:
   ```bash
   docker run -p 5001:5001 book-api
   ```
3. Verify the API is running:
   ```bash
   curl http://localhost:5001/books
   ```

### Frontend Docker Deployment

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Build the frontend Docker image:
   ```bash
   docker build -t book-frontend .
   ```
3. Run the container:
   ```bash
   docker run -p 3000:3000 book-frontend
   ```
4. Open http://localhost:3000 in your browser

## Backend Implementation Details

### Search Optimization

The search endpoint uses goroutines and channels to parallelize the search process:

1. The search function breaks the book collection into smaller chunks
2. Each chunk is processed by a separate goroutine
3. Results from each goroutine are sent to a channel
4. The main function collects all results from the channel

This approach improves search performance, especially on larger datasets.

### Storage Options

- **File Storage**: By default, the application uses a JSON file (`books.json`) for data persistence
- **MongoDB**: Pass the `-mongodb` flag to use MongoDB instead of file storage

## Frontend Implementation Details

The frontend is built with React 18 and Material-UI components, providing a modern and responsive user interface.

- **Component Structure**: Modular components for books, forms, and UI elements
- **State Management**: React Hooks for efficient state management
- **API Integration**: Axios for API requests with proper error handling
- **Styling**: Material-UI for consistent styling and responsive design
- **Validation**: Client-side form validation for data integrity

## Kubernetes Deployment

Kubernetes manifests are provided in the `kubernetes` directory for deploying the application to a Kubernetes cluster.

### Deployment Steps

1. Make sure you have kubectl and a Kubernetes cluster (or Minikube) set up
2. Deploy the backend:
   ```bash
   kubectl apply -f kubernetes/deployment.yaml
   kubectl apply -f kubernetes/service.yaml
   ```
3. Check the deployment status:
   ```bash
   kubectl get deployments
   kubectl get pods
   kubectl get services
   ```
4. Access the API:
   ```bash
   # If using Minikube
   minikube service book-api-service
   
   # If using a standard cluster
   kubectl port-forward service/book-api-service 5001:5001
   ```

## Troubleshooting

### Backend Issues

1. **API not responding**: Ensure the Go server is running and check the port number
   ```bash
   # Check if the process is running
   ps aux | grep "go run main.go"
   
   # Check if the port is in use
   netstat -ano | findstr :5001
   ```

2. **MongoDB connection failing**: Verify your MongoDB connection string and ensure MongoDB is running
   ```bash
   # Test MongoDB connection
   mongosh "mongodb+srv://yourconnectionstring"
   ```

### Frontend Issues

1. **Frontend not starting**: Check for Node.js errors
   ```bash
   cd frontend
   npm start
   ```

2. **API connection issues**: Verify the proxy setting in package.json and ensure the backend is running
   ```json
   "proxy": "http://localhost:5001"
   ```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 