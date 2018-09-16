package url_test

import (
	"reflect"
	"testing"

	"piktr/url"
)

func TestRandomizing(t *testing.T) {
	url := url.URL{}

	test := url.Generate()

	want := reflect.TypeOf("").Kind()

	got := reflect.TypeOf(test).Kind()

	if got != want {
		t.Error("Should be string")
	}

	t.Log(test)
}

func BenchmarkRandomizing(b *testing.B) {
	url := url.URL{}
	for n := 0; n < b.N; n++ {
		url.Generate()
	}
}
