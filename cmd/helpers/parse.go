package helpers

import (
	"strings"
)

func ParseInput(input string) []string {
	_, argstr, _ := strings.Cut(input, " ")

	var singleQuote bool
	var doubleQuote bool
	var backslash bool
	var arg string
	var args []string

	for _, r := range argstr {
		switch r {
		case '\'':
			if backslash && doubleQuote {
				arg += "\\"
			}
			if backslash || doubleQuote {
				arg += string(r)
			} else {
				singleQuote = !singleQuote
			}
			backslash = false

		case '"':
			if backslash || singleQuote {
				arg += string(r)
			} else {
				doubleQuote = !doubleQuote
			}
			backslash = false

		case '\\':
			if backslash || singleQuote {
				arg += string(r)
				backslash = false
			} else {
				backslash = true
			}

		case ' ':
			if backslash && doubleQuote {
				arg += "\\"
			}

			if backslash || singleQuote || doubleQuote {
				arg += string(r)
			} else if arg != "" {
				args = append(args, arg)
				arg = ""
			}
			backslash = false

		default:
			if doubleQuote && backslash {
				arg += "\\"
			}

			arg += string(r)
			backslash = false
		}
	}
	if arg != "" {
		args = append(args, arg)
	}

	return args
}

func TokenizeInput(input string) []string {
	args := make([]string, 2)

	var delim byte = ' '

	var stopIndex, continueIndex int

	var delimpresent bool

	for i := range input {

		if i == 0 {

			if input[i] == '\'' || input[i] == '"' {

				delim = input[i]

			}

			continue

		}

		if input[i] == delim {

			switch delim {

			case ' ':

				stopIndex = i

				continueIndex = i + 1

			default:

				stopIndex = i + 1

				continueIndex = stopIndex + 1

			}

			delimpresent = true

			break

		}

	}
	if !delimpresent {

		stopIndex = len(input)

		continueIndex = stopIndex

	}

	args[0] = input[:stopIndex]

	args[1] = input[continueIndex:]

	return args
}
