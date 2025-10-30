package main

import (
	"errors"
	"fmt"
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
	r.mExitCode = exitCode
	r.mRunning = false

	return nil
}

func echoExecution(r *REPL, args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments provided")
	}

	fmt.Println(strings.Join(args, " "))
	return nil
}
