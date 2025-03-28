import React, { useEffect, useState } from 'react';
import {
  Container,
  Grid,
  Typography,
  Box,
  Alert,
  Snackbar,
  CircularProgress
} from '@mui/material';
import BookCard from '../components/BookCard';
import BookForm from '../components/BookForm';
import Navbar from '../components/Navbar';
import ConfirmDialog from '../components/ConfirmDialog';
import api from '../services/api';

const HomePage = () => {
  const [books, setBooks] = useState([]);
  const [filteredBooks, setFilteredBooks] = useState([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [formOpen, setFormOpen] = useState(false);
  const [selectedBook, setSelectedBook] = useState(null);
  const [confirmDialogOpen, setConfirmDialogOpen] = useState(false);
  const [bookToDelete, setBookToDelete] = useState(null);
  const [notification, setNotification] = useState({ open: false, message: '', severity: 'success' });

  // Fetch all books on component mount
  useEffect(() => {
    fetchBooks();
  }, []);

  // Filter books when search query changes
  useEffect(() => {
    if (searchQuery) {
      handleSearch(searchQuery);
    } else {
      setFilteredBooks(books);
    }
  }, [searchQuery, books]);

  // Fetch all books from API
  const fetchBooks = async () => {
    try {
      setLoading(true);
      const data = await api.getBooks();
      setBooks(data);
      setFilteredBooks(data);
      setLoading(false);
    } catch (err) {
      setError('Failed to fetch books. Please try again later.');
      setLoading(false);
      console.error('Error fetching books:', err);
    }
  };

  // Handle search query
  const handleSearch = async (query) => {
    setSearchQuery(query);
    
    if (!query) {
      setFilteredBooks(books);
      return;
    }
    
    if (query.length < 2) return;
    
    try {
      setLoading(true);
      const results = await api.searchBooks(query);
      setFilteredBooks(results);
      setLoading(false);
    } catch (err) {
      setNotification({
        open: true,
        message: 'Error searching books',
        severity: 'error'
      });
      setLoading(false);
    }
  };

  // Open form for adding a new book
  const handleAddBook = () => {
    setSelectedBook(null);
    setFormOpen(true);
  };

  // Open form for editing a book
  const handleEditBook = (book) => {
    setSelectedBook(book);
    setFormOpen(true);
  };

  // Handle form submission (create or update)
  const handleFormSubmit = async (bookData) => {
    try {
      if (selectedBook) {
        // Update existing book
        await api.updateBook(selectedBook.bookId, bookData);
        setNotification({
          open: true,
          message: 'Book updated successfully',
          severity: 'success'
        });
      } else {
        // Create new book
        await api.createBook(bookData);
        setNotification({
          open: true,
          message: 'Book added successfully',
          severity: 'success'
        });
      }
      // Refresh books after create/update
      fetchBooks();
    } catch (err) {
      setNotification({
        open: true,
        message: selectedBook ? 'Failed to update book' : 'Failed to add book',
        severity: 'error'
      });
    }
  };

  // Open confirmation dialog before deleting
  const handleDeleteClick = (bookId) => {
    setBookToDelete(bookId);
    setConfirmDialogOpen(true);
  };

  // Delete book after confirmation
  const handleDeleteConfirm = async () => {
    try {
      await api.deleteBook(bookToDelete);
      setNotification({
        open: true,
        message: 'Book deleted successfully',
        severity: 'success'
      });
      // Refresh books after delete
      fetchBooks();
    } catch (err) {
      setNotification({
        open: true,
        message: 'Failed to delete book',
        severity: 'error'
      });
    }
    setConfirmDialogOpen(false);
    setBookToDelete(null);
  };

  // Close notification
  const handleNotificationClose = () => {
    setNotification({ ...notification, open: false });
  };

  return (
    <>
      <Navbar onAddBook={handleAddBook} onSearch={handleSearch} />
      <Container sx={{ py: 4 }} maxWidth="lg">
        {/* Page Title */}
        <Typography variant="h4" component="h1" gutterBottom>
          {searchQuery ? `Search Results: ${searchQuery}` : 'Book Collection'}
        </Typography>

        {/* Error State */}
        {error && (
          <Alert severity="error" sx={{ my: 2 }}>
            {error}
          </Alert>
        )}

        {/* Loading State */}
        {loading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}>
            <CircularProgress />
          </Box>
        ) : (
          <>
            {/* No Results State */}
            {filteredBooks.length === 0 ? (
              <Box sx={{ my: 4, textAlign: 'center' }}>
                <Typography variant="h6">
                  {searchQuery ? 'No books found matching your search.' : 'No books in the collection yet.'}
                </Typography>
                <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                  {searchQuery ? 'Try a different search term.' : 'Add some books to get started.'}
                </Typography>
              </Box>
            ) : (
              /* Books Grid */
              <Grid container spacing={4}>
                {filteredBooks.map((book) => (
                  <Grid item key={book.bookId} xs={12} sm={6} md={4}>
                    <BookCard
                      book={book}
                      onEdit={() => handleEditBook(book)}
                      onDelete={() => handleDeleteClick(book.bookId)}
                    />
                  </Grid>
                ))}
              </Grid>
            )}
          </>
        )}
      </Container>

      {/* Book Form Dialog */}
      <BookForm
        open={formOpen}
        handleClose={() => setFormOpen(false)}
        book={selectedBook}
        onSubmit={handleFormSubmit}
      />

      {/* Confirm Delete Dialog */}
      <ConfirmDialog
        open={confirmDialogOpen}
        title="Confirm Delete"
        content="Are you sure you want to delete this book? This action cannot be undone."
        onConfirm={handleDeleteConfirm}
        onCancel={() => setConfirmDialogOpen(false)}
      />

      {/* Notification Snackbar */}
      <Snackbar
        open={notification.open}
        autoHideDuration={6000}
        onClose={handleNotificationClose}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
      >
        <Alert
          onClose={handleNotificationClose}
          severity={notification.severity}
          variant="filled"
          sx={{ width: '100%' }}
        >
          {notification.message}
        </Alert>
      </Snackbar>
    </>
  );
};

export default HomePage; 