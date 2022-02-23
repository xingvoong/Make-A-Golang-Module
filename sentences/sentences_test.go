package sentences

import (
	"regexp"
	"testing"
)

//calls sentence.Sentence with a book, checking for a valid return value
func TestSentenceBook(t *testing.T) {
	name := "Swimming in the Dark"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Sentence("Swimming in the Dark")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Sentence("Swimming in the Dark") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

//calls Sentence with an empty string
func TestSentenceEmpty(t *testing.T) {
	msg, err := Sentence("")
	if msg != "" || err == nil {
		t.Fatalf(`Sentence("") = %q, %v, want "", error`, msg, err)
	}
}
