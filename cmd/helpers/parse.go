package helpers

import "strings"

func ParseInput(input string) []string {
	var result []string
	var current strings.Builder
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(input); i++ {
		char := input[i]

		switch char {
		case '\'':
			if inDoubleQuote {
				current.WriteByte(char) // Inside double quotes, treat as normal
			} else {
				inSingleQuote = !inSingleQuote // Toggle single quote state
				if !inSingleQuote {            // If closing single quote
					result = append(result, current.String())
					current.Reset()
				}
			}
		case '"':
			if inSingleQuote {
				current.WriteByte(char) // Inside single quotes, treat as normal
			} else {
				inDoubleQuote = !inDoubleQuote // Toggle double quote state
				if !inDoubleQuote {            // If closing double quote
					result = append(result, current.String())
					current.Reset()
				}
			}
		case '\\':
			if inDoubleQuote {
				// Handle escape sequences within double quotes
				if i+1 < len(input) {
					nextChar := input[i+1]
					switch nextChar {
					case '"', '\\', '$':
						current.WriteByte(nextChar) // Add the escaped character
						i++                         // Skip the next character
					case 'n':
						current.WriteString("\n") // Handle newline escape sequence
						i++                       // Skip 'n'
					default:
						current.WriteByte(char) // Just add the backslash if not followed by special char
					}
				} else {
					current.WriteByte(char) // Add the backslash if it's at the end
				}
			} else {
				current.WriteByte(char) // Outside quotes, treat as normal
			}
		case ' ':
			if inSingleQuote || inDoubleQuote {
				current.WriteByte(char) // Inside quotes, keep spaces
			} else if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset() // Reset for the next word
			}
		default:
			current.WriteByte(char) // Normal character
		}
	}

	// Add any remaining text after the loop
	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result

}
