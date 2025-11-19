package main

import (
	"strings"
)

type state int

const (
	normal state = iota
	singleQuote
)

type userCommand struct {
	input   string
	command string
	args    []string
	s       state
}

func newUserCommand(input string) *userCommand {

	c := userCommand{
		input: input,
		s:     normal,
	}
	c.parse()

	return &c
}

func (c *userCommand) parse() {
	tokens := make([]string, 0)

	working := strings.Builder{}

	for _, current := range c.input {
		switch c.s {
		case normal:
			switch current {
			case ' ':
				if working.Len() == 0 {
					continue
				}
				tokens = append(tokens, working.String())
				working.Reset()
				continue
			case '\'':
				c.s = singleQuote
				continue
			default:
				working.WriteRune(current)
			}
		case singleQuote:
			switch current {
			case '\'':
				c.s = normal
				continue
			default:
				working.WriteRune(current)
			}
		}
	}

	if working.Len() != 0 {
		tokens = append(tokens, working.String())
	}

	switch len(tokens) {
	case 0:
		c.command = ""
	case 1:
		c.command = tokens[0]
		c.args = make([]string, 0)
	default:
		c.command = tokens[0]
		c.args = tokens[1:]

	}

}
