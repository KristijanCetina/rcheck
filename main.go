package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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

func parseFlags(args []string) (*Config, error) {
	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)

	var config Config
	flagSet.StringVar(&config.Pattern, "e", "", "Regular expression pattern")
	flagSet.StringVar(&config.Input, "s", "", "Input string to match")
	flagSet.BoolVar(&config.Verbose, "v", false, "Print time spent checking")

	err := flagSet.Parse(args[1:])
	if err != nil {
		return nil, err
	}

	if config.Pattern == "" || config.Input == "" {
		return nil, fmt.Errorf("both -e (pattern) and -s (input) flags are required")
	}

	return &config, nil
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
		fmt.Println("Usage: go run program.go -e <regex_pattern> -s <input_string>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := runChecker(config); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
