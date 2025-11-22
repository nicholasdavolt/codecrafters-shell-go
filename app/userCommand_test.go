package main

import (
	"reflect"
	"testing"
)

func TestUserCommand_Parse(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantCommand    string
		wantArgs       []string
		wantFinalState state
	}{
		{
			name:           "empty string",
			input:          "",
			wantCommand:    "",
			wantArgs:       nil,
			wantFinalState: normal,
		},
		{
			name:           "single command no args",
			input:          "echo",
			wantCommand:    "echo",
			wantArgs:       nil,
			wantFinalState: normal,
		},
		{
			name:           "command with one arg",
			input:          "echo hello",
			wantCommand:    "echo",
			wantArgs:       []string{"hello"},
			wantFinalState: normal,
		},
		{
			name:           "command with multiple args",
			input:          "echo hello world",
			wantCommand:    "echo",
			wantArgs:       []string{"hello", "world"},
			wantFinalState: normal,
		},
		{
			name:           "command with leading spaces",
			input:          "  echo hello",
			wantCommand:    "echo",
			wantArgs:       []string{"hello"},
			wantFinalState: normal,
		},
		{
			name:           "command with trailing spaces",
			input:          "echo hello  ",
			wantCommand:    "echo",
			wantArgs:       []string{"hello"},
			wantFinalState: normal,
		},
		{
			name:           "command with multiple spaces between args",
			input:          "echo  hello   world",
			wantCommand:    "echo",
			wantArgs:       []string{"hello", "world"},
			wantFinalState: normal,
		},
		{
			name:           "only spaces",
			input:          "   ",
			wantCommand:    "",
			wantArgs:       nil,
			wantFinalState: normal,
		},
		{
			name:           "adjacent single quoted strings",
			input:          "echo 'hello''world'",
			wantCommand:    "echo",
			wantArgs:       []string{"helloworld"},
			wantFinalState: normal,
		},
		{
			name:           "empty single quotes",
			input:          "echo hello''world",
			wantCommand:    "echo",
			wantArgs:       []string{"helloworld"},
			wantFinalState: normal,
		},
		{
			name:           "basic double quotes",
			input:          "echo \"hello world\"",
			wantCommand:    "echo",
			wantArgs:       []string{"hello world"},
			wantFinalState: normal,
		},
		{
			name:           "double quotes multiple spaces",
			input:          "echo \"hello     world\"",
			wantCommand:    "echo",
			wantArgs:       []string{"hello     world"},
			wantFinalState: normal,
		},
		{
			name:           "double quotes surrounding single quotes",
			input:          "echo \"hello wor'ld\"",
			wantCommand:    "echo",
			wantArgs:       []string{"hello wor'ld"},
			wantFinalState: normal,
		},
		{
			name:           "basic unquoated escape",
			input:          "echo hello\\ \\ \\ \\ \\ world",
			wantCommand:    "echo",
			wantArgs:       []string{"hello     world"},
			wantFinalState: normal,
		}}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			c, err := newUserCommand(tt.input)

			if err != nil {
				t.Errorf("command = %q has returned %e", c.command, err)
				t.FailNow()
			}

			if c.command != tt.wantCommand {
				t.Errorf("command = %q, want %q", c.command, tt.wantCommand)
			}

			if !reflect.DeepEqual(c.args, tt.wantArgs) {
				t.Errorf("args = %v, want %v", c.args, tt.wantArgs)
			}

			if c.s != tt.wantFinalState {
				t.Errorf("final state = %v, want %v", c.s, tt.wantFinalState)
			}
		})
	}
}
