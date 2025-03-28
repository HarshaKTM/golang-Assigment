package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harshakumara/book-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SampleBooks returns a slice of sample book data
func SampleBooks() []models.Book {
	return []models.Book{
		{
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
		},
		{
			AuthorID:        "a7f3d45b-c892-4b7e-9f8a-2d6c4e3b5c0d",
			PublisherID:     "3e9d2f1a-b6c7-4d5e-8f9a-1b2c3d4e5f6a",
			Title:           "To Kill a Mockingbird",
			PublicationDate: "1960-07-11",
			ISBN:            "9780061120084",
			Pages:           281,
			Genre:           "Southern Gothic",
			Description:     "A story of racial injustice and moral growth, set in the American South during the 1930s.",
			Price:           12.99,
			Quantity:        10,
		},
		{
			AuthorID:        "b8e2c1d9-a7f3-4b5c-9d8e-2c1b3a4d5e6f",
			PublisherID:     "4f5e6d7c-8b9a-1c2d-3e4f-5a6b7c8d9e0f",
			Title:           "1984",
			PublicationDate: "1949-06-08",
			ISBN:            "9780451524935",
			Pages:           328,
			Genre:           "Dystopian",
			Description:     "A dystopian novel set in a totalitarian regime where critical thought is suppressed.",
			Price:           10.99,
			Quantity:        8,
		},
		{
			AuthorID:        "c9d8e7f6-5a4b-3c2d-1e0f-9a8b7c6d5e4f",
			PublisherID:     "5a6b7c8d-9e0f-1a2b-3c4d-5e6f7a8b9c0d",
			Title:           "Harry Potter and the Philosopher's Stone",
			PublicationDate: "1997-06-26",
			ISBN:            "9780590353427",
			Pages:           223,
			Genre:           "Fantasy",
			Description:     "The first book in the Harry Potter series, following the life of a young wizard and his friends.",
			Price:           20.99,
			Quantity:        15,
		},
		{
			AuthorID:        "d0e9f8a7-b6c5-4d3e-2f1a-0b9c8d7e6f5a",
			PublisherID:     "6b7c8d9e-0f1a-2b3c-4d5e-6f7a8b9c0d1e",
			Title:           "The Lord of the Rings",
			PublicationDate: "1954-07-29",
			ISBN:            "9780618640157",
			Pages:           1178,
			Genre:           "Fantasy",
			Description:     "An epic high-fantasy novel set in Middle-earth, following a hobbit's quest to destroy a powerful ring.",
			Price:           24.99,
			Quantity:        7,
		},
		{
			AuthorID:        "e1f0a9b8-c7d6-5e4f-3a2b-1c0d9e8f7a6b",
			PublisherID:     "7c8d9e0f-1a2b-3c4d-5e6f-7a8b9c0d1e2f",
			Title:           "Pride and Prejudice",
			PublicationDate: "1813-01-28",
			ISBN:            "9780141439518",
			Pages:           432,
			Genre:           "Romance",
			Description:     "A romantic novel following the character development of Elizabeth Bennet in 19th century England.",
			Price:           9.99,
			Quantity:        12,
		},
		{
			AuthorID:        "f2a1b0c9-d8e7-6f5a-4b3c-2d1e0f9a8b7c",
			PublisherID:     "8d9e0f1a-2b3c-4d5e-6f7a-8b9c0d1e2f3a",
			Title:           "The Hobbit",
			PublicationDate: "1937-09-21",
			ISBN:            "9780547928227",
			Pages:           310,
			Genre:           "Fantasy",
			Description:     "A children's fantasy novel about the adventures of a hobbit who embarks on a quest to win a share of a dragon's treasure.",
			Price:           14.99,
			Quantity:        9,
		},
		{
			AuthorID:        "a3b2c1d0-e9f8-7a6b-5c4d-3e2f1a0b9c8d",
			PublisherID:     "9e0f1a2b-3c4d-5e6f-7a8b-9c0d1e2f3a4b",
			Title:           "The Catcher in the Rye",
			PublicationDate: "1951-07-16",
			ISBN:            "9780316769488",
			Pages:           234,
			Genre:           "Coming-of-age",
			Description:     "A novel about a teenager's experiences in New York City, dealing with themes of identity, alienation, and connection.",
			Price:           11.99,
			Quantity:        6,
		},
		{
			AuthorID:        "b4c3d2e1-f0a9-8b7c-6d5e-4f3a2b1c0d9e",
			PublisherID:     "0f1a2b3c-4d5e-6f7a-8b9c-0d1e2f3a4b5c",
			Title:           "Brave New World",
			PublicationDate: "1932-10-27",
			ISBN:            "9780060850524",
			Pages:           288,
			Genre:           "Dystopian",
			Description:     "A dystopian novel set in a futuristic World State, inhabited by genetically modified citizens categorized into intelligence-based social hierarchies.",
			Price:           13.99,
			Quantity:        4,
		},
		{
			AuthorID:        "c5d4e3f2-a1b0-9c8d-7e6f-5a4b3c2d1e0f",
			PublisherID:     "1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d",
			Title:           "Moby-Dick",
			PublicationDate: "1851-10-18",
			ISBN:            "9780142437247",
			Pages:           720,
			Genre:           "Adventure",
			Description:     "The story of Captain Ahab's obsessive quest for revenge against the white whale, Moby Dick, that bit off his leg.",
			Price:           16.99,
			Quantity:        3,
		},
	}
}

// SeedBooks adds sample books to the API via HTTP requests
func SeedBooks(apiURL string) error {
	books := SampleBooks()

	fmt.Println("Seeding database with sample books...")

	for _, book := range books {
		// Generate ID if needed
		if book.ID.IsZero() {
			book.ID = primitive.NewObjectID()
		}

		// Convert book to JSON
		bookJSON, err := json.Marshal(book)
		if err != nil {
			return fmt.Errorf("error marshaling book: %v", err)
		}

		// Send POST request
		resp, err := http.Post(apiURL+"/books", "application/json", bytes.NewBuffer(bookJSON))
		if err != nil {
			return fmt.Errorf("error sending POST request: %v", err)
		}
		defer resp.Body.Close()

		// Check response status
		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		fmt.Printf("Added book: %s\n", book.Title)
	}

	fmt.Println("Sample data seeding completed successfully!")
	return nil
}
