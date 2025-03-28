package models

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBookSerialization(t *testing.T) {
	// Create a sample book
	bookID, _ := primitive.ObjectIDFromHex("bb329a31-6b1e-4daa-87ee-71631aa05866")
	book := Book{
		ID:              bookID,
		AuthorID:        "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
		PublisherID:     "2f7b19e9-b268-4440-a15b-bed8177ed607",
		Title:           "The Great Gatsby",
		PublicationDate: "1925-04-10",
		ISBN:            "9780743273565",
		Pages:           180,
		Genre:           "Novel",
		Description:     "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
		Price:           15.99,
		Quantity:        5,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Failed to marshal book to JSON: %v", err)
	}

	// Unmarshal from JSON
	var unmarshaledBook Book
	if err := json.Unmarshal(jsonData, &unmarshaledBook); err != nil {
		t.Fatalf("Failed to unmarshal book from JSON: %v", err)
	}

	// Verify fields
	if book.ID != unmarshaledBook.ID {
		t.Errorf("Expected ID %v, got %v", book.ID, unmarshaledBook.ID)
	}
	if book.Title != unmarshaledBook.Title {
		t.Errorf("Expected Title %s, got %s", book.Title, unmarshaledBook.Title)
	}
	if book.AuthorID != unmarshaledBook.AuthorID {
		t.Errorf("Expected AuthorID %s, got %s", book.AuthorID, unmarshaledBook.AuthorID)
	}
	if book.PublisherID != unmarshaledBook.PublisherID {
		t.Errorf("Expected PublisherID %s, got %s", book.PublisherID, unmarshaledBook.PublisherID)
	}
	if book.PublicationDate != unmarshaledBook.PublicationDate {
		t.Errorf("Expected PublicationDate %s, got %s", book.PublicationDate, unmarshaledBook.PublicationDate)
	}
	if book.ISBN != unmarshaledBook.ISBN {
		t.Errorf("Expected ISBN %s, got %s", book.ISBN, unmarshaledBook.ISBN)
	}
	if book.Pages != unmarshaledBook.Pages {
		t.Errorf("Expected Pages %d, got %d", book.Pages, unmarshaledBook.Pages)
	}
	if book.Genre != unmarshaledBook.Genre {
		t.Errorf("Expected Genre %s, got %s", book.Genre, unmarshaledBook.Genre)
	}
	if book.Description != unmarshaledBook.Description {
		t.Errorf("Expected Description %s, got %s", book.Description, unmarshaledBook.Description)
	}
	if book.Price != unmarshaledBook.Price {
		t.Errorf("Expected Price %f, got %f", book.Price, unmarshaledBook.Price)
	}
	if book.Quantity != unmarshaledBook.Quantity {
		t.Errorf("Expected Quantity %d, got %d", book.Quantity, unmarshaledBook.Quantity)
	}
}
