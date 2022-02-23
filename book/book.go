package main

import (
	"fmt"

	"example.com/sentences"
)

func main() {
	// Get a recommendation sentence and print it
	message := sentences.Sentence("Ultralearning")
	fmt.Println(message)
}
