package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Ale"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Ale")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ale") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
