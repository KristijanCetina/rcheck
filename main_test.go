package main

import (
	"testing"
)

func TestCheckRegexMatch(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
		wantErr bool
	}{
		{"^hello", "hello world", true, false},
		{"world$", "hello world", true, false},
		{"[0-9]+", "abc123", true, false},
		{"[a-z]+", "123", false, false},
		{"[", "invalid", false, true}, // Invalid regex pattern
	}

	for _, tt := range tests {
		got, err := checkRegexMatch(tt.pattern, tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("checkRegexMatch(%q, %q) error = %v, wantErr %v", tt.pattern, tt.input, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("checkRegexMatch(%q, %q) = %v, want %v", tt.pattern, tt.input, got, tt.want)
		}
	}
}
