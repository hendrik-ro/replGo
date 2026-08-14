package repl

import "testing"

func TestChangePrompt(t *testing.T) {
	c := DefaultConfig()
	c.ChangePrompt("Your input")
	if c.Prompt != "Your input " {
		t.Errorf("expected prompt to be 'Your input ' but got '%s'", c.Prompt)
	}
}
