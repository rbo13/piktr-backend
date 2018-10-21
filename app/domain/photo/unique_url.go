// Package photo handles generating of UniqueURL field in photos table
package photo

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// GenerateUniqueURL returns a generated unique url in the pattern of AdjectiveNounAnimal.
func GenerateUniqueURL() string {

	adjective := randomize(Adjectives())
	noun := randomize(Nouns())
	animal := randomize(Animals())
	return formatURL(strings.Join([]string{adjective, noun, animal}, " "))
}

// randomize a given slice of words
func randomize(words []string) string {
	// seed the initial randomization
	rand.Seed(time.Now().Unix() * len(words))
	return words[rand.Intn(len(words))]
}

// formats a given url
func formatURL(url string) string {
	url = strings.Title(url)
	return removeWhitespace(url)
}

// effeciently remove white space in a string
func removeWhitespace(url string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, url)
}
