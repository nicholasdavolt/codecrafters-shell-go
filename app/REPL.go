package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type REPL struct {
	mRunning  bool
	mExitCode int
	mCommands []Command
}

type Command struct {
	name string
	desc string
	exec func(r *REPL, args []string) error
}

func newREPL() *REPL {

	return &REPL{
		mRunning: false,
	}
}
func exitExecution(r *REPL, args []string) error {
	if len(args) == 0 {

		return errors.New("no exit code provided")
	}
	exitCode, err := strconv.Atoi(args[0])

	if err != nil {

		return err
	}
	r.mExitCode = exitCode
	r.mRunning = false

	return nil
}

func (r *REPL) start() {
	exitCommand := Command{
		name: "exit",
		desc: "exit command, takes an exit code",
		exec: exitExecution,
	}

	r.mCommands = append(r.mCommands, exitCommand)
	r.mRunning = true

	r.read()

}

func (r *REPL) read() {
	scanner := bufio.NewScanner(os.Stdin)
	for r.mRunning {
		fmt.Fprint(os.Stdout, "$ ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		r.ealuate(input)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func (r *REPL) ealuate(input string) {
	uC := newUserCommand(input)

	for _, cmd := range r.mCommands {
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
