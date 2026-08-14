package repl

import "testing"

func TestHelpCommand(t *testing.T) {
	r := newREPL()
	if len(r.Config.Commands) != 2 {
		t.Errorf("expected 2 commands, got %d", len(r.Config.Commands))
	}
	if r.Config.Commands[0].Name != "help" {
		t.Errorf("expected command name 'help', got '%s'", r.Config.Commands[0].Name)
	}
	if r.Config.Commands[0].Description != "Display this help message" {
		t.Errorf("expected command description 'Display this help message', got '%s'", r.Config.Commands[0].Description)
	}
	if r.Config.Commands[0].Handler == nil {
		t.Errorf("expected command handler, got nil")
	}

	if r.Config.Commands[1].Name != "exit" {
		t.Errorf("expected command name 'exit', got '%s'", r.Config.Commands[1].Name)
	}
	if r.Config.Commands[1].Description != "Exit the REPL" {
		t.Errorf("expected command description 'Exit the REPL', got '%s'", r.Config.Commands[1].Description)
	}
	if r.Config.Commands[1].Handler == nil {
		t.Errorf("expected command handler, got nil")
	}
}
