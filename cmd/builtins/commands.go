package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var commands = map[string]bool{
	"exit": true, "type": true, "echo": true, "ls": true,
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
		command := exec.Command(strings.Split(input, " ")[0], strings.Split(input, " ")[1:]...)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Printf("%s: command not found\n", strings.Split(input, " ")[0])
		}
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
	fmt.Printf("%s: not found\n", input[5:len(input)-1])
}

func echoHandler(input string) {
	fmt.Printf("%s\n", input[5:len(input)-1])
}

func exitHandler(input string) {
	if input == "exit 0\n" {
		os.Exit(0)
	}
}
