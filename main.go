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

func main() {
	// Define named flags
	expressionPtr := flag.String("e", "", "Regular expression pattern")
	stringPtr := flag.String("s", "", "Input string to match against")
	flagVerbose := flag.Bool("v", false, "Print time spent checking")

	// Parse command-line flags
	flag.Parse()

	// Check if both flags are provided
	if *expressionPtr == "" || *stringPtr == "" {
		fmt.Println("Usage: go run program.go -e <regex_pattern> -s <input_string>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Record start time
	start := time.Now()
	// Check for match
	matched, err := checkRegexMatch(*expressionPtr, *stringPtr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Calculate duration
	duration := time.Since(start)
	// Print result
	if matched {
		fmt.Println("\033[32mThe pattern matches the input string!\033[0m")
	} else {
		fmt.Println("\033[0;31mThe pattern does not match the input string.\033[0m")
	}
	// Print time spent
	if *flagVerbose {
		fmt.Printf("Time spent checking match: %v\n", duration)
	}
}
