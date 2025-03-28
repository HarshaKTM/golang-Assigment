package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/harshakumara/book-api/config"
	"github.com/harshakumara/book-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if using MongoDB or file storage
	if config.MongoClient != nil {
		// MongoDB implementation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := config.BookCollection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		var books []models.Book
		if err = cursor.All(ctx, &books); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(books)
	} else {
		// File storage implementation
		fs := config.NewFileStorage("books.json")
		books, err := fs.ReadBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(books)
	}
}

// GetBook returns a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	if config.MongoClient != nil {
		// MongoDB implementation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var book models.Book
		err = config.BookCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&book)
		if err != nil {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
	} else {
		// File storage implementation
		fs := config.NewFileStorage("books.json")
		books, err := fs.ReadBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var book models.Book
		found := false
		for _, b := range books {
			if b.ID.Hex() == id {
				book = b
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

// CreateBook creates a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate new ID if not provided
	if book.ID.IsZero() {
		book.ID = primitive.NewObjectID()
	}

	if config.MongoClient != nil {
		// MongoDB implementation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := config.BookCollection.InsertOne(ctx, book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// File storage implementation
		fs := config.NewFileStorage("books.json")
		books, err := fs.ReadBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books = append(books, book)
		if err := fs.WriteBooks(books); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if config.MongoClient != nil {
		// MongoDB implementation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		// Preserve the original ID
		updatedBook.ID = objID

		_, err = config.BookCollection.ReplaceOne(ctx, bson.M{"_id": objID}, updatedBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// File storage implementation
		fs := config.NewFileStorage("books.json")
		books, err := fs.ReadBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		found := false
		for i, b := range books {
			if b.ID.Hex() == id {
				// Preserve the original ID
				updatedBook.ID = b.ID
				books[i] = updatedBook
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		if err := fs.WriteBooks(books); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(updatedBook)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	if config.MongoClient != nil {
		// MongoDB implementation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		_, err = config.BookCollection.DeleteOne(ctx, bson.M{"_id": objID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// File storage implementation
		fs := config.NewFileStorage("books.json")
		books, err := fs.ReadBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		found := false
		filteredBooks := []models.Book{}
		for _, b := range books {
			if b.ID.Hex() != id {
				filteredBooks = append(filteredBooks, b)
			} else {
				found = true
			}
		}

		if !found {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		if err := fs.WriteBooks(filteredBooks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// SearchBooks searches for books based on title and description
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the search keyword from query parameters
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		http.Error(w, "Search keyword is required", http.StatusBadRequest)
		return
	}

	keyword = strings.ToLower(keyword)

	var allBooks []models.Book
	var err error

	if config.MongoClient != nil {
		// MongoDB implementation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Using $or to search in both title and description with case-insensitive search
		filter := bson.M{
			"$or": []bson.M{
				{"title": bson.M{"$regex": keyword, "$options": "i"}},
				{"description": bson.M{"$regex": keyword, "$options": "i"}},
			},
		}

		cursor, err := config.BookCollection.Find(ctx, filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &allBooks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// File storage implementation
		fs := config.NewFileStorage("books.json")
		allBooks, err = fs.ReadBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Use goroutines and channels for concurrent search
		results := make(chan models.Book)
		var wg sync.WaitGroup

		// Divide the books into chunks for parallel processing
		chunkSize := 10 // Adjust based on expected dataset size
		if len(allBooks) < chunkSize {
			chunkSize = 1
		}

		chunks := (len(allBooks) + chunkSize - 1) / chunkSize

		// Launch goroutines to search each chunk
		for i := 0; i < chunks; i++ {
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()

				for j := start; j < end && j < len(allBooks); j++ {
					book := allBooks[j]
					title := strings.ToLower(book.Title)
					description := strings.ToLower(book.Description)

					if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
						results <- book
					}
				}
			}(i*chunkSize, (i+1)*chunkSize)
		}

		// Close the channel once all goroutines are done
		go func() {
			wg.Wait()
			close(results)
		}()

		// Collect results
		var searchResults []models.Book
		for book := range results {
			searchResults = append(searchResults, book)
		}

		allBooks = searchResults
	}

	json.NewEncoder(w).Encode(allBooks)
}
