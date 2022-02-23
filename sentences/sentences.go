package sentences

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Sentence returns a recommendation sentence for the named book.
func Sentence(book_name string) (string, error) {
	// If no book name was given, return an error with messsage of empty name
	if book_name == "" {
		return "", errors.New("empty name")
	}

	// return a recommendation sentence embeds the name of a book in a message
	message := fmt.Sprintf(randomFormat(), book_name)
	return message, nil
}

// Sentences returns a map that map each of the book name with a recommendation message
func Sentences(book_names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, book := range book_names {
		message, err := Sentence(book)
		if err != nil {
			return nil, err
		}
		messages[book] = message
	}

	return messages, nil
}

// init the rand function base on time
func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns one of a set of recommended message.  The message is chosed randomly.
func randomFormat() string {
	sentences := []string{
		"You would like to read %v",
		"How about you try %v",
		"I really enjoyed %v, you may too.",
	}
	return sentences[rand.Intn(len(sentences))]
}
