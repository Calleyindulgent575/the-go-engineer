package main

import "fmt"

// ============================================================================
// Section 7: Strings & Text — Log Parser (Exercise Starter)
// Level: Intermediate
// ============================================================================
//
// EXERCISE: Build a Log File Parser
//
// REQUIREMENTS:
//  1. [ ] Define a `LogEntry` struct with Timestamp, Level, and Message fields
//  2. [ ] Compile a regex pattern to match log lines like:
//         [2024-01-15 14:30:22] ERROR: Connection timeout
//  3. [ ] Implement `parseLine(line string) (LogEntry, bool)` using the regex
//  4. [ ] Read a multi-line log string using `bufio.Scanner`
//  5. [ ] Parse each line, collect entries, and print a summary
//
// HINTS:
//   - Use regexp.MustCompile to compile the pattern at startup
//   - Pattern: `\[(.+?)\] (\w+): (.+)` captures timestamp, level, message
//   - Use strings.NewReader to create an io.Reader from a string
//   - bufio.NewScanner(reader) reads line-by-line
//
// RUN: go run ./07-strings-and-text/6-log-parser/_starter
// SOLUTION: See the main.go file in the parent directory
// ============================================================================

// TODO: Define your LogEntry struct

// TODO: Compile your regex pattern

// TODO: Implement parseLine function

func main() {
	fmt.Println("=== Log Parser Exercise ===")
	fmt.Println()
	fmt.Println("TODO: Implement your log parser!")
	fmt.Println("See the REQUIREMENTS above for what to build.")
	fmt.Println()
	fmt.Println("When finished, compare your solution with ../main.go")
}
