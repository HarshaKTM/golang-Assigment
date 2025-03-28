package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SeedFileStorage adds sample books directly to the JSON file storage
func SeedFileStorage(filePath string) error {
	fmt.Println("Seeding file storage with sample books...")

	books := SampleBooks()

	// Assign IDs to all books
	for i := range books {
		if books[i].ID.IsZero() {
			books[i].ID = primitive.NewObjectID()
		}
	}

	// Convert books to JSON
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling books: %v", err)
	}

	// Write to file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Printf("Added %d sample books to file: %s\n", len(books), filePath)

	return nil
}
