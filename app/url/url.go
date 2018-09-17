package url

import (
	"math/rand"
	"strings"
	"time"
	"unicode"

	"piktr/app/url/url_helper"
)

// URL defines the set of
// URL pattern.
// Sample: /WittyFaucetSunbittern
type URL struct {
	UniqueURL string
}

// Generate generates the
// unique URL pattern which is
// AdjectiveNounAnimal
func (u *URL) Generate() string {
	adjective := randomize(url_helper.Adjectives())
	noun := randomize(url_helper.Nouns())
	animal := randomize(url_helper.Animals())
	u.UniqueURL = formatURL(strings.Join([]string{adjective, noun, animal}, " "))

	return u.UniqueURL
}

// randomize a given slice of words
func randomize(words []string) string {
	// seed the initial randomization
	rand.Seed(time.Now().Unix() * int64(len(words)))
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
