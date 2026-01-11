package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func checkRegexMatch(pattern, input string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, fmt.Errorf("invalid regex pattern: %v", err)
	}
	match := re.FindString(input)
	return match == input, nil
}

type Config struct {
	Pattern string
	Input   string
	Verbose bool
}

func readFromStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading stdin: %v", err)
	}
	return input.String(), nil
}

func parseFlags(args []string) (*Config, error) {
	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(flagSet.Output(), "Usage: %s -e <regex_pattern> [-s <input_string>]\n", args[0])
		flagSet.PrintDefaults()
	}

	var pattern, input string
	var verbose bool
	flagSet.StringVar(&pattern, "e", "", "Regular expression pattern")
	flagSet.StringVar(&input, "s", "", "Input string to match (optional, reads from stdin if not provided)")
	flagSet.BoolVar(&verbose, "v", false, "Print time spent checking")

	err := flagSet.Parse(args[1:])
	if err != nil {
		return nil, err
	}

	if pattern == "" {
		return nil, fmt.Errorf("the -e (pattern) flag is required")
	}

	if input == "" {
		input, err = readFromStdin()
		if err != nil {
			return nil, err
		}
		if input == "" {
			return nil, fmt.Errorf("no input provided: use -s flag or pipe input via stdin")
		}
	}

	return &Config{Pattern: pattern, Input: input, Verbose: verbose}, nil
}

func printResult(matched bool, duration time.Duration, verbose bool) {
	if matched {
		fmt.Println("\033[32mThe pattern matches the input string!\033[0m")
	} else {
		fmt.Println("\033[0;31mThe pattern does not match the input string.\033[0m")
	}

	if verbose {
		fmt.Printf("Time spent checking: %v\n", duration)
	}
}

func runChecker(config *Config) error {
	start := time.Now()
	matched, err := checkRegexMatch(config.Pattern, config.Input)
	if err != nil {
		return err
	}
	duration := time.Since(start)

	printResult(matched, duration, config.Verbose)
	return nil
}

func main() {
	config, err := parseFlags(os.Args)
	if err != nil {
		os.Exit(1)
	}

	if err := runChecker(config); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
