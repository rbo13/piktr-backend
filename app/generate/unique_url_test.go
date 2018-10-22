package generate_test

import (
	"piktr/app/generate"
	"testing"
)

func TestShouldGenerateUniqueURL(t *testing.T) {

	uniqueURL := generate.UniqueURL()

	if uniqueURL == "" {
		t.Errorf("Expects to be: %v, but got an empty string", uniqueURL)
	}

	t.Log(uniqueURL)
}
