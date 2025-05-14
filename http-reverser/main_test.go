package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	r := Reverse("abcd")

	if r != "dcba" {
		t.Errorf("Expected 'dcba', got '%s'", r)
	}
}

func TestArgValidator(t *testing.T) {
	tests := []struct {
		arg      string
		expected string
	}{
		{"", "arg is required"},
		{"abcd", "dcba"},
	}

	for _, test := range tests {
		t.Run(test.arg, func(t *testing.T) {
			r := ArgValidator(test.arg)
			if r != test.expected {
				t.Errorf("Expected '%s', got '%s'", test.expected, r)
			}
		})
	}
}
