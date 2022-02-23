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

	// Request a recommender message
	message, err := sentences.Sentence("Ultralearning")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
