package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/harshakumara/book-api/models"
)

// FileStorage represents a file-based data storage
type FileStorage struct {
	filePath string
	mutex    sync.RWMutex
}

// NewFileStorage creates a new instance of FileStorage
func NewFileStorage(filePath string) *FileStorage {
	// Create file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		// Initialize with empty array
		if _, err := file.Write([]byte("[]")); err != nil {
			log.Fatalf("Failed to initialize file: %v", err)
		}
	}

	return &FileStorage{
		filePath: filePath,
		mutex:    sync.RWMutex{},
	}
}

// ReadBooks reads all books from the file
func (fs *FileStorage) ReadBooks() ([]models.Book, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	data, err := ioutil.ReadFile(fs.filePath)
	if err != nil {
		return nil, err
	}

	var books []models.Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, err
	}

	return books, nil
}

// WriteBooks writes books to the file
func (fs *FileStorage) WriteBooks(books []models.Book) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fs.filePath, data, 0644)
}
