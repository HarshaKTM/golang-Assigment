package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/harshakumara/book-api/utils"
)

func main() {
	// Parse command line flags
	apiURL := flag.String("url", "http://localhost:5001", "Base URL of the Book API")
	flag.Parse()

	// Seed the database
	if err := utils.SeedBooks(*apiURL); err != nil {
		log.Fatalf("Error seeding database: %v", err)
		os.Exit(1)
	}

	fmt.Println("Database seeded successfully with sample book data!")
}
