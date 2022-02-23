package sentences

import "fmt"

// Sentence returns a recommendation sentence for the named book.
func Sentence(book_name string) string {
	// return a recommendation sentence embeds the name of a book in a message
	message := fmt.Sprintf("You would like to read %v", book_name)
	return message
}
