package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type REPL struct {
	running  bool
	exitCode int
	commands []Command
}

type Command struct {
	name string
	desc string
	exec func(r *REPL, args []string) error
}

func newREPL() *REPL {

	return &REPL{}
}

func (r *REPL) start() {

	r.running = true
	r.registerCommands()
	r.read()

}

func (r *REPL) registerCommands() {
	exitCommand := Command{
		name: "exit",
		desc: "exit command, takes an exit code",
		exec: exitExec,
	}

	echoCommand := Command{
		name: "echo",
		desc: "echo command",
		exec: echoExec,
	}

	typeCommand := Command{
		name: "type",
		desc: "type command",
		exec: typeExec,
	}
	pwdCommand := Command{
		name: "pwd",
		desc: "prints current working directory",
		exec: pwdExec,
	}
	cdCommand := Command{
		name: "cd",
		desc: "changes current working directory",
		exec: cdExec,
	}

	r.commands = append(r.commands, exitCommand, echoCommand, typeCommand, pwdCommand, cdCommand)
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

	trimmed := strings.TrimSpace(input)

	if len(trimmed) == 0 {
		return
	}

	uC, err := newUserCommand(trimmed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, cmd := range r.commands {
		if uC.command == cmd.name {
			err := cmd.exec(r, uC.args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			return
		}
	}

	_, err = exec.LookPath(uC.command)

	if err == nil {
		cmd := exec.Command(uC.command, uC.args...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		return

	}
	r.printBadCommand(uC.command)

}

func (r *REPL) printBadCommand(input string) {
	fmt.Fprintln(os.Stderr, input+": command not found")
}
