package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Cmd struct {
	command string
	args    []string
}

var commands = map[string]bool{
	"exit": true, "type": true, "echo": true, "ls": true, "path": true,
}

func BuiltinHandler(input string) {
	cmd := &Cmd{
		command: strings.Split(input, " ")[0],
		args:    strings.Split(input, " ")[1:],
	}
	switch cmd.command {
	case "exit":
		exitHandler()
	case "type":
		typeHandler(cmd.args)
	case "echo":
		echoHandler(cmd.args)
	default:
		command := exec.Command(cmd.command, cmd.args...)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Printf("%s: command not found\n", strings.Split(strings.TrimRight(input, "\n"), " ")[0])
		}
	}
}

func typeHandler(input []string) {
	_, exists := commands[input[0]]
	if exists {
		fmt.Printf("%s is a shell builtin\n", input[0])
		return
	}

	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		fp := filepath.Join(path, input[0])

		if _, err := os.Stat(fp); err == nil {
			fmt.Println(fp)
			return
		}
	}
	fmt.Printf("%s: not found\n", input[0])
}

func echoHandler(input []string) {
	fmt.Printf("%s\n", strings.Join(input, " "))
}

func exitHandler() {
	os.Exit(0)
}
