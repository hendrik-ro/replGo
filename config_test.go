package repl

import "testing"

func TestChangePrompt(t *testing.T) {
	c := DefaultConfig()
	c.ChangePrompt("Your input")
	if c.Prompt != "Your input " {
		t.Errorf("expected prompt to be 'Your input ' but got '%s'", c.Prompt)
	}
}

func TestUpdateStartup(t *testing.T) {
	c := DefaultConfig()
	c.UpdateStartup("Starting your app...")
	if c.System.Startup != "Starting your app..." {
		t.Errorf("expected startup message to be 'Starting your app...' but got '%s'", c.System.Startup)
	}
	c.UpdateStartup("")
	if c.System.Startup != DefaultSystem.Startup {
		t.Errorf("expected startup message to be '%s' but got '%s'", DefaultSystem.Startup, c.System.Startup)
	}
}

func TestUpdateShutdown(t *testing.T) {
	c := DefaultConfig()
	c.UpdateShutdown("Closing down...")
	if c.System.Shutdown != "Closing down..." {
		t.Errorf("expected shutdown message to be 'Closing down...' but got '%s'", c.System.Shutdown)
	}
	c.UpdateShutdown("")
	if c.System.Shutdown != DefaultSystem.Shutdown {
		t.Errorf("expected shutdown message to be '%s' but got '%s'", DefaultSystem.Shutdown, c.System.Shutdown)
	}
}
