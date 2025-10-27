package main

import (
	"bufio"
	"fmt"
	"os"
)

type REPL struct {
	mRunning bool
}

func newREPL() *REPL {
	return &REPL{
		mRunning: false,
	}
}

func (r *REPL) start() {
	r.mRunning = true
	r.read()

}

func (r *REPL) read() {
	scanner := bufio.NewScanner(os.Stdin)
	for r.mRunning {
		fmt.Fprint(os.Stdout, "$ ")
		scanner.Scan()
		input := scanner.Text()
		r.print(input)
	}
}

func (r *REPL) ealuate() {
	//TODO implement evalutaion
}

func (r *REPL) print(input string) {
	fmt.Fprintln(os.Stdout, input+": command not found")
}
