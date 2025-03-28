# Book API Frontend

A modern React frontend for the Book API project. This application provides a user-friendly interface to manage books, including creating, reading, updating, and deleting books. It also features an optimized search functionality that connects to the backend's search endpoint.

## Features

- Modern UI built with Material-UI components
- Complete CRUD operations for books
- Real-time search with optimized backend integration
- Responsive design works on all device sizes
- Form validation for data integrity
- Notification system for user feedback

## Technologies Used

- React 18 for UI components
- React Router for navigation
- Material-UI for styled components
- Axios for API requests
- React Hooks for state management

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn
- Book API backend running on port 5001

### Installation

1. Install dependencies:
   ```
   npm install
   ```
   or
   ```
   yarn install
   ```

2. Start the development server:
   ```
   npm start
   ```
   or
   ```
   yarn start
   ```

The application will be available at http://localhost:3000.

## Project Structure

- `/src/components`: Reusable UI components
- `/src/pages`: Page components
- `/src/services`: API service layer
- `/src/utils`: Utility functions
- `/src/assets`: Static assets like images

## Integration with Backend

The frontend is configured to work with the Book API backend running on port 5001. The proxy is set up in `package.json` to forward API requests to the backend.

## Form Validation

The application includes client-side validation for all book fields:
- Required fields (title, authorId, publisherId, ISBN)
- Numeric validation (pages, price, quantity)
- Date format validation (publicationDate)

## Search Functionality

The search functionality is optimized to:
- Debounce user input to limit API calls
- Filter results on the client-side for short queries
- Make API calls for more complex searches 