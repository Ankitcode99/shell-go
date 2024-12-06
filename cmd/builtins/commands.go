package builtins

import (
	"fmt"
	"os"
)

func BuiltinHandler(cmd, input string) {

	switch cmd {
	case "exit":
		exitHandler(input)
	// case "pwd":
	// 	PwdHandler()
	case "echo":
		echoHandler(input)
	default:
		fmt.Printf("%s: command not found\n", input[:len(input)-1])
	}
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
