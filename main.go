package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/harshakumara/book-api/api/handlers"
	"github.com/harshakumara/book-api/config"
	"github.com/harshakumara/book-api/utils"
)

func main() {
	// Command line flags
	useMongoDb := flag.Bool("mongodb", false, "Use MongoDB for storage instead of file")
	port := flag.String("port", "5001", "Port to run the server on")
	seedData := flag.Bool("seed", false, "Seed the database with sample data")
	flag.Parse()

	// Initialize storage
	if *useMongoDb {
		log.Println("Using MongoDB for storage")
		config.ConnectDB()

		// Seed MongoDB if flag is set
		if *seedData {
			log.Println("Seeding MongoDB with sample data...")
			// MongoDB seeding happens on /books endpoint access
		}
	} else {
		log.Println("Using file-based storage")
		// Create the storage file if it doesn't exist
		config.NewFileStorage("books.json")

		// Seed file storage if flag is set
		if *seedData {
			log.Println("Seeding file storage with sample data...")
			if err := utils.SeedFileStorage("books.json"); err != nil {
				log.Fatalf("Failed to seed file storage: %v", err)
			}
		}
	}

	// Initialize router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")

	// Set up server
	serverPort := *port
	if envPort := os.Getenv("PORT"); envPort != "" {
		serverPort = envPort
	}

	log.Printf("Server starting on port %s...\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, r))
}
