package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) {
	// Remove newline character
	line = line[:len(line)-1]

	if strings.HasPrefix(line, "exit") {
		// get exit args
		args := strings.Split(line, " ")
		if len(args) <= 1 {
			fmt.Println("exit: need 1 argument")
			return
		}

		// get exit code
		rawCode := args[1]
		res, err := strconv.Atoi(rawCode)
		if err != nil {
			fmt.Println("exit: argument must be an integer")
			return
		}
		os.Exit(res)

	}

	fmt.Println(line + ": command not found")
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		parseLine(line)
	}
}
