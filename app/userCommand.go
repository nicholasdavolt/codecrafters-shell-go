package main

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

	working := ""

	for i := 0; i < len(c.input); i++ {
		current := string(c.input[i])

		switch c.s {
		case normal:
			if current == " " {
				if working == "" {
					continue
				}
				tokens = append(tokens, working)
				working = ""
				continue
			}

			working += current
		}

	}

	if working != "" {
		tokens = append(tokens, working)
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
