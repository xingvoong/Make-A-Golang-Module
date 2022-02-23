package main

import (
	"fmt"
	"log"

	"example.com/sentences"
)

func main() {
	// Get a recommendation sentence and print it
	log.SetPrefix("sentences: ")
	log.SetFlags(0)

	// A slice of books
	books := []string{"Ultralearning", "Geek Love", "Deep Work"}

	// Request a recommender message
	messages, err := sentences.Sentences(books)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
