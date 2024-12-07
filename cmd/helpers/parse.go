package helpers

import "strings"

func ParseInput(input string) []string {
	var parts []string
	var current strings.Builder
	inQuotes := false
	quoteChar := ' ' // Track which quote character is currently open

	for _, char := range input {
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
		default:
			current.WriteRune(char) // Add character to current part
		}
	}

	// Add the last part if it exists
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}
