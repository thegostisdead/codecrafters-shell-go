package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLine(line string) {
	// Remove newline character
	line = line[:len(line)-1]
	fmt.Println(line + ": command not found")
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	parseLine(line)
}
