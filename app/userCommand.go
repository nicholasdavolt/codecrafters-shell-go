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

const escapes = `\"`

type userCommand struct {
	input   string
	command string
	args    []string
	s       state
	ps      state
}

func newUserCommand(input string) (*userCommand, error) {

	c := userCommand{
		input: input,
		s:     normal,
		ps:    normal,
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
	runeInput := []rune(c.input)

	for i, current := range runeInput {
		switch c.s {
		case normal:
			switch current {
			case ' ':
				if working.Len() == 0 {
					continue
				}
				tokens = append(tokens, working.String())
				working.Reset()

			case '\'':
				c.s = singleQuote
				c.ps = normal
			case '"':
				c.s = doubleQuote
				c.ps = normal
			case '\\':
				c.s = escape
				c.ps = normal
			default:
				working.WriteRune(current)
			}
		case singleQuote:
			switch current {
			case '\'':
				c.s = normal
				c.ps = singleQuote
			default:
				working.WriteRune(current)
			}
		case doubleQuote:
			switch current {
			case '"':
				c.s = normal
				c.ps = doubleQuote

			case '\\':

				if isEscape(runeInput[i+1]) {
					c.ps = doubleQuote
					c.s = escape

				} else {
					working.WriteRune(current)
				}

			default:
				working.WriteRune(current)

			}
		case escape:

			working.WriteRune(current)

			c.s = c.ps
			c.ps = escape
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

func isEscape(r rune) bool {
	for _, b := range escapes {
		if r == b {
			return true
		}
	}
	return false
}
