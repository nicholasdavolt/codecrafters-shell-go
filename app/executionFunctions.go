package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func exitExec(r *REPL, args []string) error {
	if len(args) == 0 {
		r.running = false
		return nil
	}
	exitCode, err := strconv.Atoi(args[0])

	if err != nil {

		return err
	}
	r.exitCode = exitCode
	r.running = false

	return nil
}

func echoExec(r *REPL, args []string) error {

	fmt.Println(strings.Join(args, " "))
	return nil
}

func pwdExec(r *REPL, args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(cwd)

	return nil
}

func cdExec(r *REPL, args []string) error {
	if len(args) != 1 {
		return errors.New("Not enough arguments")
	}

	path := args[0]
	err := os.Chdir(path)

	if err != nil {
		return errors.New("cd: " + path + ": no such file or directory")
	}

	return nil
}

func typeExec(r *REPL, args []string) error {
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
