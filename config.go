package repl

func DefaultConfig() *Config {
	return &Config{
		Prompt:   "> ",
		Commands: DefaultCommands,
		History: &History{
			MaxSize: 100,
			Entries: &[]HistoryEntry{},
		},
	}
}

func (c *Config) ChangePrompt(prompt string) {
	// Update prompt with a trailing space to separate input from prompt
	c.Prompt = prompt + " "
}
