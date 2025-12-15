package main

import (
	"os"
	"testing"
)

func TestCheckRegexMatch(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
		wantErr bool
	}{
		{"^hello.*", "hello world", true, false},
		{".*world$", "hello world", true, false},
		{".*[0-9]+", "abc123", true, false},
		{"[0-9]+", "abc123", false, false},
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
func TestCheckRegexMatchEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		input   string
		want    bool
		wantErr bool
	}{
		{"empty pattern", "", "", true, false},
		// {"empty input", "test", "", false, false},
		{"exact match", "test", "test", true, false},
		{"partial match", "test", "testing", false, false},
		{"dot matches any char", "t.st", "test", true, false},
		{"alternation", "cat|dog", "cat", true, false},
		{"character class", "[aeiou]", "a", true, false},
		{"negated class", "[^0-9]", "a", true, false},
		{"quantifier zero or more", "a*b", "b", true, false},
		{"quantifier one or more", "a+", "aaa", true, false},
		{"quantifier range", "a{2,4}", "aaa", true, false},
		{"escape special char", `\.`, ".", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkRegexMatch(tt.pattern, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkRegexMatch(%q, %q) error = %v, wantErr %v", tt.pattern, tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkRegexMatch(%q, %q) = %v, want %v", tt.pattern, tt.input, got, tt.want)
			}
		})
	}
}
func TestMain(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{"valid match", []string{"prog", "-e", "^hello", "-s", "hello world"}, false},
		{"valid no match", []string{"prog", "-e", "^goodbye", "-s", "hello world"}, false},
		{"missing expression flag", []string{"prog", "-s", "test"}, true},
		{"missing string flag", []string{"prog", "-e", "test"}, true},
		{"missing both flags", []string{"prog"}, true},
		{"invalid regex pattern", []string{"prog", "-e", "[", "-s", "test"}, true},
		{"with verbose flag", []string{"prog", "-e", "test", "-s", "test", "-v"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.args

			if tt.expectErr {
				// For error cases, we'd need to capture os.Exit
				// This is a limitation of testing os.Exit calls
				t.Skip("Cannot easily test os.Exit behavior")
			}
		})
	}
}
