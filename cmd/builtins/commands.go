package builtins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var commands = map[string]bool{
	"exit": true,
	"type": true,
	"echo": true,
	"ls":   true,
}

func BuiltinHandler(cmd, input string) {

	switch cmd {
	case "exit":
		exitHandler(input)
	case "type":
		typeHandler(input)
	case "echo":
		echoHandler(input)
	default:
		fmt.Printf("%s: command not found\n", input[:len(input)-1])
	}
}

func typeHandler(input string) {

	_, exists := commands[input[5:len(input)-1]]
	if exists {
		fmt.Printf("%s is a shell builtin\n", input[5:len(input)-1])
		return
	}

	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		fp := filepath.Join(path, input[5:len(input)-1])

		if _, err := os.Stat(fp); err == nil {

			fmt.Println(fp)

			return

		}
	}
	fmt.Printf("%s not found\n", input[5:len(input)-1])
}

func echoHandler(input string) {

	fmt.Printf("%s\n", input[5:len(input)-1])
}

func exitHandler(input string) {
	// fmt.Print("Got in exitHandler\n")
	if input == "exit 0\n" {
		// fmt.Printf("Goodbye!!")
		os.Exit(0)
	}
}
