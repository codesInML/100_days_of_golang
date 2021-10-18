package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ifeoluwa")

	got := buffer.String()
	want := "Hello, Ifeoluwa"

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}
