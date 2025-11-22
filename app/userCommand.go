package main

import (
	"fmt"
	"strings"
)

type state int

const (
	normal state = iota
	singleQuote
	doubleQuote
	escape
)

type userCommand struct {
	input   string
	command string
	args    []string
	s       state
}

func newUserCommand(input string) (*userCommand, error) {

	c := userCommand{
		input: input,
		s:     normal,
	}
	err := c.parse()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *userCommand) parse() error {
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
			case '"':
				c.s = doubleQuote
				continue
			case '\\':
				c.s = escape
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
		case doubleQuote:
			switch current {
			case '"':
				c.s = normal
				continue
			default:
				working.WriteRune(current)
			}
		case escape:
			working.WriteRune(current)
			c.s = normal
		}

	}

	if c.s != normal {
		return fmt.Errorf("improper quoting")
	}

	if working.Len() != 0 {
		tokens = append(tokens, working.String())
	}

	switch len(tokens) {
	case 0:
		c.command = ""
	case 1:
		c.command = tokens[0]
		c.args = nil
	default:
		c.command = tokens[0]
		c.args = tokens[1:]

	}

	return nil
}
