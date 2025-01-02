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
func TestHellos(t *testing.T) {
	names := []string{"Ale", "Aarón", "Hugo"}
	want := map[string]*regexp.Regexp{
		"Ale":   regexp.MustCompile(`\bAle\b`),
		"Aarón": regexp.MustCompile(`\bAarón\b`),
		"Hugo":  regexp.MustCompile(`\bHugo\b`),
	}

	messages, err := Hellos(names)

	if err != nil {
		t.Fatalf(`Hellos(%v) returned error: %v`, names, err)
	}

	for name, msg := range messages {
		if !want[name].MatchString(msg) {
			t.Fatalf(`Hellos(%v) = %q, want match for %#q`, name, msg, want[name])
		}
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"Ale", "", "Aarón"}

	_, err := Hellos(names)

	if err == nil {
		t.Fatalf(`Hellos(%v) = nil, want error`, names)
	}
}
