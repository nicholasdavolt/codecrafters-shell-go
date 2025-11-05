package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type REPL struct {
	running  bool
	exitCode int
	path     []string
	commands []Command
}

type Command struct {
	name string
	desc string
	exec func(r *REPL, args []string) error
}

func newREPL() *REPL {

	return &REPL{
		running: false,
	}
}

func (r *REPL) start() {

	r.running = true
	r.registerCommands()
	r.registerPath()

	r.read()

}

func (r *REPL) registerCommands() {
	exitCommand := Command{
		name: "exit",
		desc: "exit command, takes an exit code",
		exec: exitExecution,
	}

	echoCommand := Command{
		name: "echo",
		desc: "echo command",
		exec: echoExecution,
	}

	typeCommand := Command{
		name: "type",
		desc: "type command",
		exec: typeExcecution,
	}

	r.commands = append(r.commands, exitCommand, echoCommand, typeCommand)
}

func (r *REPL) registerPath() {
	osPath := os.Getenv("PATH")

	if osPath == "" {
		fmt.Fprintln(os.Stderr, "PATH environment variable not set")
		return
	}

	sep := string(os.PathListSeparator)

	r.path = append(r.path, strings.Split(osPath, sep)...)
}

func (r *REPL) read() {
	scanner := bufio.NewScanner(os.Stdin)
	for r.running {
		fmt.Fprint(os.Stdout, "$ ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		r.evaluate(input)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func (r *REPL) evaluate(input string) {
	uC := newUserCommand(input)

	for _, cmd := range r.commands {
		if uC.mCommand == cmd.name {
			err := cmd.exec(r, uC.mArgs)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			return
		}
	}
	r.printBadCommand(input)

}

func (r *REPL) printBadCommand(input string) {
	fmt.Fprintln(os.Stdout, input+": command not found")
}
