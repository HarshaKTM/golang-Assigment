import axios from 'axios';

const API_URL = '/books'; // Using proxy in package.json

const api = {
  // Fetch all books
  getBooks: async () => {
    try {
      const response = await axios.get(API_URL);
      return response.data;
    } catch (error) {
      console.error('Error fetching books:', error);
      throw error;
    }
  },

  // Fetch a book by ID
  getBook: async (id) => {
    try {
      const response = await axios.get(`${API_URL}/${id}`);
      return response.data;
    } catch (error) {
      console.error(`Error fetching book with ID ${id}:`, error);
      throw error;
    }
  },

  // Create a new book
  createBook: async (bookData) => {
    try {
      const response = await axios.post(API_URL, bookData);
      return response.data;
    } catch (error) {
      console.error('Error creating book:', error);
      throw error;
    }
  },

  // Update a book
  updateBook: async (id, bookData) => {
    try {
      const response = await axios.put(`${API_URL}/${id}`, bookData);
      return response.data;
    } catch (error) {
      console.error(`Error updating book with ID ${id}:`, error);
      throw error;
    }
  },

  // Delete a book
  deleteBook: async (id) => {
    try {
      await axios.delete(`${API_URL}/${id}`);
      return true;
    } catch (error) {
      console.error(`Error deleting book with ID ${id}:`, error);
      throw error;
    }
  },

  // Search books
  searchBooks: async (query) => {
    try {
      const response = await axios.get(`${API_URL}/search?q=${encodeURIComponent(query)}`);
      return response.data;
    } catch (error) {
      console.error(`Error searching books with query "${query}":`, error);
      throw error;
    }
  }
};

export default api; 