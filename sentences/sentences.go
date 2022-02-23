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

// init the rand function base on time
func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns one of a set of recommended message.  The message is chosed randomly.
func randomFormat() string {
	sentences := []string{
		"You would like to read %v",
		"How about yoy try %v",
		"I really enjoyed %v, you may too.",
	}
	return sentences[rand.Intn(len(sentences))]
}
