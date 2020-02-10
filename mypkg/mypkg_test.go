package mypkg

import (
	"testing"
)

func TestContains(t *testing.T) {
	got := Contains("This is a sentence", "is a")
	want := true
	if got != want {
		t.Errorf("Contains: got %v, but wanted %v", got, want)
	}

	got = Contains("This is a sentence", "what")
	want = false
	if got != want {
		t.Errorf("Not contains: got %v, but wanted %v", got, want)
	}

	got = Contains("This is a sentence", "")
	want = true
	if got != want {
		t.Errorf("Empty str2: got %v, but wanted %v", got, want)
	}
}
