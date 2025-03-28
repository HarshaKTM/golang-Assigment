import React from 'react';
import { 
  Card, 
  CardContent, 
  CardActions, 
  Typography, 
  Button, 
  Chip,
  Stack,
  Box
} from '@mui/material';
import { 
  Edit as EditIcon,
  Delete as DeleteIcon, 
  MenuBook as MenuBookIcon
} from '@mui/icons-material';

const BookCard = ({ book, onEdit, onDelete }) => {
  return (
    <Card 
      sx={{ 
        height: '100%', 
        display: 'flex', 
        flexDirection: 'column',
        transition: 'all 0.3s',
        '&:hover': {
          transform: 'translateY(-5px)',
          boxShadow: 6
        }
      }}
    >
      <CardContent sx={{ flexGrow: 1 }}>
        <Typography gutterBottom variant="h5" component="h2" fontWeight="bold">
          {book.title}
        </Typography>
        
        <Stack direction="row" spacing={1} my={1}>
          <Chip 
            label={book.genre} 
            size="small" 
            color="primary" 
            variant="outlined"
            icon={<MenuBookIcon />} 
          />
          <Chip 
            label={`$${book.price.toFixed(2)}`} 
            size="small" 
            color="success" 
            variant="outlined" 
          />
        </Stack>
        
        <Typography variant="body2" color="text.secondary" paragraph>
          {book.description.length > 150 
            ? `${book.description.substring(0, 150)}...` 
            : book.description}
        </Typography>
        
        <Box sx={{ mt: 2 }}>
          <Typography variant="body2" color="text.secondary">
            <strong>ISBN:</strong> {book.isbn}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            <strong>Pages:</strong> {book.pages}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            <strong>Published:</strong> {book.publicationDate}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            <strong>In Stock:</strong> {book.quantity}
          </Typography>
        </Box>
      </CardContent>
      
      <CardActions>
        <Button 
          size="small" 
          startIcon={<EditIcon />} 
          onClick={() => onEdit(book)}
        >
          Edit
        </Button>
        <Button 
          size="small" 
          color="error" 
          startIcon={<DeleteIcon />} 
          onClick={() => onDelete(book.bookId)}
        >
          Delete
        </Button>
      </CardActions>
    </Card>
  );
};

export default BookCard; 