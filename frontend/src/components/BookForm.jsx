import React, { useState, useEffect } from 'react';
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
  Grid,
  InputAdornment
} from '@mui/material';

const initialFormState = {
  title: '',
  authorId: '',
  publisherId: '',
  publicationDate: '',
  isbn: '',
  pages: '',
  genre: '',
  description: '',
  price: '',
  quantity: ''
};

const BookForm = ({ open, handleClose, book, onSubmit }) => {
  const [formData, setFormData] = useState(initialFormState);
  const [errors, setErrors] = useState({});

  useEffect(() => {
    if (book) {
      setFormData({
        title: book.title || '',
        authorId: book.authorId || '',
        publisherId: book.publisherId || '',
        publicationDate: book.publicationDate || '',
        isbn: book.isbn || '',
        pages: book.pages || '',
        genre: book.genre || '',
        description: book.description || '',
        price: book.price || '',
        quantity: book.quantity || ''
      });
    } else {
      setFormData(initialFormState);
    }
    setErrors({});
  }, [book, open]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    
    // Validate numeric fields
    if (name === 'pages' || name === 'quantity') {
      if (value && !/^[0-9]+$/.test(value)) {
        return;
      }
    } else if (name === 'price') {
      if (value && !/^[0-9]*\.?[0-9]*$/.test(value)) {
        return;
      }
    }
    
    setFormData({
      ...formData,
      [name]: value
    });
    
    // Clear error when field is being edited
    if (errors[name]) {
      setErrors({
        ...errors,
        [name]: null
      });
    }
  };

  const validateForm = () => {
    const newErrors = {};
    
    // Required fields
    if (!formData.title) newErrors.title = 'Title is required';
    if (!formData.authorId) newErrors.authorId = 'Author ID is required';
    if (!formData.publisherId) newErrors.publisherId = 'Publisher ID is required';
    if (!formData.isbn) newErrors.isbn = 'ISBN is required';
    
    // Format validations
    if (formData.publicationDate && !/^\d{4}-\d{2}-\d{2}$/.test(formData.publicationDate)) {
      newErrors.publicationDate = 'Use format YYYY-MM-DD';
    }
    
    // Numeric validations
    if (formData.pages && isNaN(Number(formData.pages))) {
      newErrors.pages = 'Pages must be a number';
    }
    
    if (formData.price && isNaN(Number(formData.price))) {
      newErrors.price = 'Price must be a number';
    }
    
    if (formData.quantity && isNaN(Number(formData.quantity))) {
      newErrors.quantity = 'Quantity must be a number';
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validateForm()) {
      // Convert numeric fields
      const processedData = {
        ...formData,
        pages: formData.pages ? parseInt(formData.pages, 10) : 0,
        price: formData.price ? parseFloat(formData.price) : 0,
        quantity: formData.quantity ? parseInt(formData.quantity, 10) : 0
      };
      
      // Pass bookId from existing book if editing
      if (book && book.bookId) {
        processedData.bookId = book.bookId;
      }
      
      onSubmit(processedData);
      handleClose();
    }
  };

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="md" fullWidth>
      <DialogTitle>{book ? 'Edit Book' : 'Add New Book'}</DialogTitle>
      <DialogContent>
        <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 2 }}>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                label="Title"
                name="title"
                value={formData.title}
                onChange={handleChange}
                error={!!errors.title}
                helperText={errors.title}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                label="Genre"
                name="genre"
                value={formData.genre}
                onChange={handleChange}
                error={!!errors.genre}
                helperText={errors.genre}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                label="Author ID"
                name="authorId"
                value={formData.authorId}
                onChange={handleChange}
                error={!!errors.authorId}
                helperText={errors.authorId}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                label="Publisher ID"
                name="publisherId"
                value={formData.publisherId}
                onChange={handleChange}
                error={!!errors.publisherId}
                helperText={errors.publisherId}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                label="ISBN"
                name="isbn"
                value={formData.isbn}
                onChange={handleChange}
                error={!!errors.isbn}
                helperText={errors.isbn}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                fullWidth
                label="Publication Date (YYYY-MM-DD)"
                name="publicationDate"
                value={formData.publicationDate}
                onChange={handleChange}
                error={!!errors.publicationDate}
                helperText={errors.publicationDate}
              />
            </Grid>
            <Grid item xs={12} sm={4}>
              <TextField
                fullWidth
                label="Pages"
                name="pages"
                type="number"
                value={formData.pages}
                onChange={handleChange}
                error={!!errors.pages}
                helperText={errors.pages}
              />
            </Grid>
            <Grid item xs={12} sm={4}>
              <TextField
                fullWidth
                label="Price"
                name="price"
                value={formData.price}
                onChange={handleChange}
                error={!!errors.price}
                helperText={errors.price}
                InputProps={{
                  startAdornment: <InputAdornment position="start">$</InputAdornment>,
                }}
              />
            </Grid>
            <Grid item xs={12} sm={4}>
              <TextField
                fullWidth
                label="Quantity"
                name="quantity"
                type="number"
                value={formData.quantity}
                onChange={handleChange}
                error={!!errors.quantity}
                helperText={errors.quantity}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Description"
                name="description"
                multiline
                rows={4}
                value={formData.description}
                onChange={handleChange}
                error={!!errors.description}
                helperText={errors.description}
              />
            </Grid>
          </Grid>
        </Box>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Cancel</Button>
        <Button onClick={handleSubmit} variant="contained" color="primary">
          {book ? 'Update' : 'Add'} Book
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default BookForm; 