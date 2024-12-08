package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	handler "github.com/codecrafters-io/shell-starter-go/cmd/builtins"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		handler.BuiltinHandler(input)
	}

}
