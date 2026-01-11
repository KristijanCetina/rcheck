package main

import (
	"testing"
	"time"
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
		{"escaped special character", `\.`, ".", true, false},
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
func TestParseFlags(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    *Config
		wantErr bool
	}{
		{"valid match", []string{"prog", "-e", "^hello", "-s", "hello world"}, &Config{Pattern: "^hello", Input: "hello world", Verbose: false}, false},
		{"valid no match", []string{"prog", "-e", "^goodbye", "-s", "hello world"}, &Config{Pattern: "^goodbye", Input: "hello world", Verbose: false}, false},
		{"with verbose flag", []string{"prog", "-e", "test", "-s", "test", "-v"}, &Config{Pattern: "test", Input: "test", Verbose: true}, false},
		{"missing expression flag", []string{"prog", "-s", "test"}, nil, true},
		{"missing string flag", []string{"prog", "-e", "test"}, nil, true},
		{"missing both flags", []string{"prog"}, nil, true},
		{"invalid flag", []string{"prog", "-x", "invalid"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFlags(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("parseFlags() returned nil, expected config")
					return
				}
				if got.Pattern != tt.want.Pattern || got.Input != tt.want.Input || got.Verbose != tt.want.Verbose {
					t.Errorf("parseFlags() = %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestRunChecker(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{"valid match", &Config{Pattern: "^hello", Input: "hello world", Verbose: false}, false},
		{"valid no match", &Config{Pattern: "^goodbye", Input: "hello world", Verbose: false}, false},
		{"invalid regex pattern", &Config{Pattern: "[", Input: "test", Verbose: false}, true},
		{"with verbose", &Config{Pattern: "test", Input: "test", Verbose: true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runChecker(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("runChecker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrintResult(t *testing.T) {
	tests := []struct {
		name     string
		matched  bool
		duration time.Duration
		verbose  bool
	}{
		{"match no verbose", true, time.Millisecond, false},
		{"no match no verbose", false, time.Millisecond, false},
		{"match with verbose", true, time.Millisecond, true},
		{"no match with verbose", false, time.Millisecond, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test just ensures the function doesn't panic
			// In a real scenario, you might capture stdout to verify output
			printResult(tt.matched, tt.duration, tt.verbose)
		})
	}
}
