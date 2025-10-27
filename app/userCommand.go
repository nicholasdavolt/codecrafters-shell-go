package main

import (
	"fmt"
)

type state int

const (
	normal state = iota
	singleQote
)

type userCommand struct {
	mInput   string
	mCommand string
	mArgs    []string
	mState   state
}

func newUserCommand(input string) *userCommand {
	c := userCommand{
		mInput: input,
		mState: normal,
	}
	c.parse()
	return &c
}

func (c *userCommand) parse() {
	tokens := make([]string, 0)
	working := ""
	for i := 0; i < len(c.mInput); i++ {
		current := string(c.mInput[i])

		switch c.mState {
		case normal:
			if current == " " {
				tokens = append(tokens, working)
				working = ""
				continue
			}

			working += current
		}

	}
	tokens = append(tokens, working)

	switch len(tokens) {
	case 0:
		fmt.Println("You need to provide a command.")
	case 1:
		c.mCommand = tokens[0]
	default:
		c.mCommand = tokens[0]
		c.mArgs = tokens[1:]

	}
}
