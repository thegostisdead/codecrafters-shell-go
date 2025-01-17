package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// shell built-in commands
var builtInCommands = []string{"echo", "exit", "type", "pwd", "cd"}

func doEcho(args []string) {
	if len(args) <= 1 {
		fmt.Println("echo: need 1 argument")
		return
	}
	fmt.Println(strings.Join(args[1:], " "))
	return
}
func doExit(args []string) {
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
func doType(args []string) {
	if len(args) <= 1 {
		fmt.Println("type: need 1 argument")
		return
	}
	cmd := args[1]
	if contains(builtInCommands, cmd) {
		fmt.Println(cmd + " is a shell builtin")
	} else if searchBinInPath(cmd) != "" {
		fmt.Println(cmd + " is " + searchBinInPath(cmd))
	} else {
		fmt.Println(cmd + ": not found")
	}
}
func searchBinInPath(cmd string) string {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	for _, p := range paths {
		if _, err := os.Stat(p + "/" + cmd); err == nil {
			return p + "/" + cmd
		}

	}
	return ""
}
func contains(commands []string, cmd string) bool {

	for _, c := range commands {
		if c == cmd {
			return true
		}
	}
	return false
}

func doPwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(dir)
}

func doCd(args []string) {

	if len(args) <= 1 {
		// go to home directory
		err := os.Chdir(os.Getenv("HOME"))
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	dir := args[1]

	if strings.HasPrefix(dir, "~") {
		// replace ~ with home directory
		dir = os.Getenv("HOME") + strings.Replace(args[1], "~", "/", -1)
	}

	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("cd: " + dir + ": No such file or directory")
	}
}

func parseLine(line string) {
	line = line[:len(line)-1]
	args := strings.Split(line, " ")

	switch args[0] {
	case "echo":
		doEcho(args)
	case "exit":
		doExit(args)
	case "type":
		doType(args)
	case "pwd":
		doPwd()
	case "cd":
		doCd(args)
	default:
		binPath := searchBinInPath(args[0])
		if binPath != "" {
			command := exec.Command(args[0], args[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", args[0])
				return
			}

		} else {
			fmt.Println(line + ": command not found")
		}
	}

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
