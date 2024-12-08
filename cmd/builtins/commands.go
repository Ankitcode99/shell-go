package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	helper "github.com/codecrafters-io/shell-starter-go/cmd/helpers"
)

type Cmd struct {
	command string
	args    []string
}

var commands = map[string]bool{
	"exit": true, "type": true, "echo": true, "ls": true, "path": true, "pwd": true, "cd": true, // "cat": true,
}

func BuiltinHandler(input string) {
	helperRes := helper.TokenizeInput(input)

	cmd := Cmd{
		command: string(helperRes[0]),
		args:    helperRes[1:], //helper.ParseInput(input)[1:],
	}

	// fmt.Printf("%#v\n", cmd)

	switch cmd.command {
	case "exit":
		exitHandler()
	case "type":
		typeHandler(cmd.args)
	case "echo":
		echoHandler(helper.ParseInput(input))
	case "pwd":
		pwdHandler()
	case "cd":
		cdHandler(cmd.args)
	case "cat":
		catHandler(input)
	default:
		if input[0] == '\'' || input[0] == '"' {
			// fmt.Printf("INSIDE CUSTOM EXEC\n")
			ch := input[0]
			st := 0
			for i := 1; i < len(input); i++ {
				if (input[i] == '\'' || input[i] == '"') && ch == input[i] {
					st = i + 1
					break
				}
			}

			newInput := "cat " + input[st+1:]
			// fmt.Printf("%s\n", newInput)
			catHandler(newInput)
		} else {
			command := exec.Command(cmd.command, cmd.args...)

			command.Stdout = os.Stdout
			command.Stderr = os.Stderr

			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", strings.Split(strings.TrimRight(input, "\n"), " ")[0])
			}
		}
	}
}

func catHandler(input string) {
	parsedInput := helper.ParseInput(input)[1:]
	// fmt.Printf("Parsed Input: %#v\n", parsedInput)
	for _, filePath := range parsedInput {
		fileContent, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Print("Something went wrong")
		}
		fmt.Print(string(fileContent))

		// fmt.Printf("%s\n", filePath)
	}
}

func cdHandler(input []string) {
	if input[0] == "~" {
		os.Chdir(os.Getenv("HOME"))
		return
	}

	p := path.Clean(input[0])

	if !path.IsAbs(p) {

		dir, _ := os.Getwd()

		p = path.Join(dir, p)

	}

	err := os.Chdir(p)
	if err != nil {

		fmt.Printf("cd: %s: No such file or directory\n", p)

	}
}

func pwdHandler() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
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
	var inputs []string
	if input[0][0] == '\'' {
		for i := 0; i < len(input); i++ {
			input[i] = strings.Trim(input[i], "'")
		}
		fmt.Printf("%s\n", strings.Join(input, " "))
	} else {
		for i := 0; i < len(input); i++ {
			input[i] = strings.Trim(input[i], "'")
			// fmt.Printf("%d - %s - %d\n", i, input[i], len(input[i]))
			if len(input[i]) > 0 {
				// fmt.Print("\nRemoving space!\n")
				inputs = append(inputs, input[i])
			}
		}
		fmt.Printf("%s\n", strings.Join(inputs, " "))
	}

}

func exitHandler() {
	os.Exit(0)
}
