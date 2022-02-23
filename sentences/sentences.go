package sentences

import (
	"errors"
	"fmt"
)

// Sentence returns a recommendation sentence for the named book.
func Sentence(book_name string) (string, error) {
	// If no book name was given, return an error with messsage of empty name
	if book_name == "" {
		return "", errors.New("empty name")
	}

	// return a recommendation sentence embeds the name of a book in a message
	message := fmt.Sprintf("You would like to read %v", book_name)
	return message, nil
}
