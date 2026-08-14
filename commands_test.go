package repl

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	r := newREPL(Config{
		Commands: []Command{},
	})

	cmd := Command{
		Name:        "test",
		Description: "tests the add method",
		Handler: func(args []string) {
			fmt.Println(args)
		},
	}
	r.Add(cmd)

	if len(r.Config.Commands) != 1 {
		t.Errorf("expected 1 command, got %d", len(r.Config.Commands))
	}

	if r.Config.Commands[0].Name != "test" {
		t.Errorf("expected command name 'test', got '%s'", r.Config.Commands[0].Name)
	}
	if r.Config.Commands[0].Description != "tests the add method" {
		t.Errorf("expected command description 'tests the add method', got '%s'", r.Config.Commands[0].Description)
	}
	if r.Config.Commands[0].Handler == nil {
		t.Errorf("expected command handler, got nil")
	}
}
