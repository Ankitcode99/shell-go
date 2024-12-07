package helpers

import "strings"

func ParseInput(input string) []string {
	var parts []string
	var current strings.Builder
	inQuotes := false
	quoteChar := ' ' // Track which quote character is currently open

	for i := 0; i < len(input); i++ {
		char := rune(input[i])
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
				current.WriteRune(char)
			} else {
				if i+1 < len(input) {
					current.WriteRune(rune(input[i+1]))
					i++
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
