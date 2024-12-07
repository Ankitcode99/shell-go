package helpers

import (
	"fmt"

	"github.com/buildkite/shellwords"
)

func ParseInput(input string) []string {
	// var parts []string
	// var current strings.Builder
	// inQuotes := false
	// // startQuoteChar := ' ' // Track which quote character is currently open
	// escapeNext := false // Flag to indicate if the next character should be treated as escaped
	// quotes := []rune{}

	// for i := 0; i < len(input); i++ {
	// 	char := rune(input[i])
	// 	// fmt.Printf("%d - %t - %v\n", i, inQuotes, quoteChar)
	// 	switch char {
	// 	case '"':
	// 		end := -1
	// 		for j := i + 1; j < len(input); j++ {
	// 			if input[j] == '"' {
	// 				end = j
	// 				break
	// 			}
	// 		}

	// 		s := "hello'shell'\\n'script"

	// 		fmt.Println("Compare - ", input[i+1:end])

	// 		v := fmt.Sprintf("%s", s)
	// 		x := fmt.Sprintf("%s", input[i+1:end])
	// 		fmt.Printf("s = %s\n", s)
	// 		fmt.Printf("v = %s\n", v)
	// 		fmt.Printf("x = %s\n", x)
	// 		fmt.Printf("Double quote  - %s\n", input[i+1:end])

	// 		parts = append(parts, fmt.Sprintf("%s", input[i+1:end]))
	// 		i = end

	// 	case '\'':

	// 		currStr := ""

	// 		for j := i + 1; j < len(input); j++ {
	// 			if input[j] == '\'' {
	// 				currStr += string(input[i+1 : j])
	// 				fmt.Printf("single quote  - %s\n", input[i+1:j])
	// 				parts = append(parts, currStr)
	// 				i = j
	// 				break
	// 			}
	// 		}

	// 	case ' ':
	// 		if inQuotes {
	// 			current.WriteRune(char) // Keep spaces inside quotes
	// 		} else {
	// 			if current.Len() > 0 {
	// 				parts = append(parts, current.String())
	// 				current.Reset()
	// 			}
	// 		}
	// 	case '\\':
	// 		if inQuotes {
	// 			// fmt.Printf(" INSIDE inQuotes \n")
	// 			if i+1 < len(input) && (input[i+1] == '\\' || input[i+1] == '$' || input[i+1] == '\n' || input[i+1] == '\'' || input[i+1] == '"') {
	// 				// fmt.Printf(" INSIDE ESCAPE \n")
	// 				// Write the next character as is (escaped)
	// 				if quotes[0] == '"' {
	// 					current.WriteRune(rune(input[i+1]))
	// 					i++
	// 				} else {
	// 					current.WriteRune(rune(input[i]))
	// 				}
	// 			} else {
	// 				current.WriteRune(char)
	// 			}
	// 			escapeNext = true // Set flag to escape the next character within quotes
	// 		} else {
	// 			// fmt.Printf("%d - %c\n", i, char)
	// 			if i+1 < len(input) && (input[i+1] == '\\' || input[i+1] == '$' || input[i+1] == '\n' || input[i+1] == '\'' || input[i+1] == '"') {
	// 				// fmt.Printf(" INSIDE ESCAPE \n")
	// 				current.WriteRune(rune(input[i+1])) // Write the next character as is (escaped)
	// 				i++                                 // Skip the next character since it's escaped
	// 			} else {
	// 				current.WriteRune(char)
	// 			}
	// 		}
	// 	default:
	// 		if escapeNext {
	// 			current.WriteRune(char) // Add the escaped character directly
	// 			escapeNext = false      // Reset escape flag
	// 		} else {
	// 			current.WriteRune(char) // Add regular character to current part
	// 		}
	// 	}
	// }

	// // Add the last part if it exists
	// if current.Len() > 0 {
	// 	parts = append(parts, current.String())
	// }

	words, err := shellwords.Split(input)

	if err != nil {
		fmt.Println("Error parsing command:", err)
		return []string{}
	}

	// Print the parsed words
	for _, word := range words {
		fmt.Println(word)
	}

	return words
}
