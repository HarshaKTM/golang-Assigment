package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/harshakumara/book-api/config"
	"github.com/harshakumara/book-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetBooks(t *testing.T) {
	// Initialize file storage with test data
	fs := config.NewFileStorage("../../test_books.json")

	// Create sample books
	book1 := models.Book{
		ID:              primitive.NewObjectID(),
		Title:           "Test Book 1",
		AuthorID:        "author1",
		PublisherID:     "publisher1",
		PublicationDate: "2023-01-01",
		ISBN:            "1234567890",
		Pages:           100,
		Genre:           "Test",
		Description:     "Test description 1",
		Price:           9.99,
		Quantity:        10,
	}

	book2 := models.Book{
		ID:              primitive.NewObjectID(),
		Title:           "Test Book 2",
		AuthorID:        "author2",
		PublisherID:     "publisher2",
		PublicationDate: "2023-02-01",
		ISBN:            "0987654321",
		Pages:           200,
		Genre:           "Test",
		Description:     "Test description 2",
		Price:           19.99,
		Quantity:        5,
	}

	// Write test books to file
	testBooks := []models.Book{book1, book2}
	_ = fs.WriteBooks(testBooks)

	// Create a test request
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(GetBooks)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var responseBooks []models.Book
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBooks); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	// Verify response
	if len(responseBooks) != 2 {
		t.Errorf("Expected 2 books, got %d", len(responseBooks))
	}
}

func TestCreateBook(t *testing.T) {
	// Initialize file storage
	fs := config.NewFileStorage("../../test_books.json")

	// Empty books array
	_ = fs.WriteBooks([]models.Book{})

	// Create a test book
	newBook := models.Book{
		Title:           "New Test Book",
		AuthorID:        "author1",
		PublisherID:     "publisher1",
		PublicationDate: "2023-03-01",
		ISBN:            "1122334455",
		Pages:           150,
		Genre:           "Test",
		Description:     "New test description",
		Price:           14.99,
		Quantity:        3,
	}

	// Convert book to JSON
	bookJSON, _ := json.Marshal(newBook)

	// Create a test request
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(CreateBook)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check if book was added
	books, _ := fs.ReadBooks()
	if len(books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(books))
	}

	// Cleanup
	_ = fs.WriteBooks([]models.Book{})
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/search", SearchBooks).Methods("GET")
	return r
}
