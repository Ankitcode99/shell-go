package helpers

import (
	"strings"
)

func ParseInput(input string) []string {
	var parts []string
	var current strings.Builder
	inQuotes := false
	quoteChar := ' '    // Track which quote character is currently open
	escapeNext := false // Flag to indicate if the next character should be treated as escaped

	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		// fmt.Printf("%d - %c\n", i, char)
		switch char {
		case '"', '\'':
			if !inQuotes {
				inQuotes = true
				quoteChar = char // Set the current quote character
			} else if char == quoteChar {
				inQuotes = false // Close the current quote
			} else {
				current.WriteRune(char) // Keep different quote types inside
			}
		case ' ':
			if inQuotes {
				current.WriteRune(char) // Keep spaces inside quotes
			} else {
				if current.Len() > 0 {
					parts = append(parts, current.String())
					current.Reset()
				}
			}
		case '\\':
			if inQuotes {
				// fmt.Printf(" INSIDE inQuotes \n")
				if i+1 < len(input) && (input[i+1] == '\\' || input[i+1] == '$' || input[i+1] == '\n' || input[i+1] == '\'' || input[i+1] == '"') {
					// fmt.Printf(" INSIDE ESCAPE \n")
					current.WriteRune(rune(input[i+1])) // Write the next character as is (escaped)
					i++                                 // Skip the next character since it's escaped
				} else {
					current.WriteRune(char)
				}
				escapeNext = true // Set flag to escape the next character within quotes
			} else {
				// fmt.Printf("%d - %c\n", i, char)
				if i+1 < len(input) && (input[i+1] == '\\' || input[i+1] == '$' || input[i+1] == '\n' || input[i+1] == '\'' || input[i+1] == '"') {
					// fmt.Printf(" INSIDE ESCAPE \n")
					current.WriteRune(rune(input[i+1])) // Write the next character as is (escaped)
					i++                                 // Skip the next character since it's escaped
				} else {
					current.WriteRune(char)
				}
			}
		default:
			if escapeNext {
				current.WriteRune(char) // Add the escaped character directly
				escapeNext = false      // Reset escape flag
			} else {
				current.WriteRune(char) // Add regular character to current part
			}
		}
	}

	// Add the last part if it exists
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}
