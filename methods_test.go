package repl

import (
	"fmt"
	"testing"
)

var numDefaultCommands = len(DefaultCommands)

func TestAdd(t *testing.T) {
	r := NewREPL()

	cmd := Command{
		Name:        "test",
		Description: "tests the add method",
		Handler: func(cfg *Config, args []string) {
			fmt.Println(args)
		},
	}
	r.Add(cmd)

	if len(r.Config.Commands) != numDefaultCommands+1 {
		t.Errorf("expected %d commands, got %d", numDefaultCommands+1, len(r.Config.Commands))
	}

	if r.Config.Commands[numDefaultCommands].Name != "test" {
		t.Errorf("expected command name 'test', got '%s'", r.Config.Commands[numDefaultCommands].Name)
	}
	if r.Config.Commands[numDefaultCommands].Description != "tests the add method" {
		t.Errorf("expected command description 'tests the add method', got '%s'", r.Config.Commands[numDefaultCommands].Description)
	}
	if r.Config.Commands[numDefaultCommands].Handler == nil {
		t.Errorf("expected command handler, got nil")
	}
}

func TestRemove(t *testing.T) {
	r := NewREPL()

	cmd := Command{
		Name:        "test",
		Description: "tests the remove method",
		Handler: func(cfg *Config, args []string) {
			fmt.Println(args)
		},
	}
	r.Add(cmd)

	if len(r.Config.Commands) != numDefaultCommands+1 {
		t.Errorf("expected %d commands, got %d", numDefaultCommands+1, len(r.Config.Commands))
	}

	r.Remove("test")

	if len(r.Config.Commands) != numDefaultCommands {
		t.Errorf("expected %d commands, got %d", numDefaultCommands, len(r.Config.Commands))
	}
}

func TestList(t *testing.T) {
	r := NewREPL()
	cmds := r.List()
	if len(*cmds) != numDefaultCommands {
		t.Errorf("expected %d commands, got %d", numDefaultCommands, len(*cmds))
	}
}
