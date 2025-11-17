package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func exitExecution(r *REPL, args []string) error {
	if len(args) == 0 {

		return errors.New("no exit code provided")
	}
	exitCode, err := strconv.Atoi(args[0])

	if err != nil {

		return err
	}
	r.exitCode = exitCode
	r.running = false

	return nil
}

func echoExecution(r *REPL, args []string) error {

	fmt.Println(strings.Join(args, " "))
	return nil
}

func typeExcecution(r *REPL, args []string) error {
	if len(args) != 1 {
		return errors.New("wrong number of arguments")
	}

	arg := args[0]

	for _, cmd := range r.commands {
		if cmd.name == arg {
			fmt.Println(arg + " is a shell builtin")
			return nil
		}
	}

	path, err := exec.LookPath(arg)

	if err != nil {
		fmt.Println(arg + ": not found")
		return nil
	}

	fmt.Println(arg + " is " + path)
	return nil
}
