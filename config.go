package repl

var DefaultPrompt = "> "
var DefaultSystem = System{
	Startup:  "Starting REPL...",
	Shutdown: "Shutting down...",
}

func DefaultConfig() *Config {
	return &Config{
		Prompt:   DefaultPrompt,
		Commands: DefaultCommands,
		History: &History{
			MaxSize: 100,
			Entries: &[]HistoryEntry{},
		},
		System: DefaultSystem,
	}
}

func (cfg *Config) ChangePrompt(prompt string) {
	// Update prompt with a trailing space to separate input from prompt
	cfg.Prompt = prompt + " "
}

func (cfg *Config) UpdateStartup(startup string) {
	// Update startup message if provided
	if startup == "" {
		cfg.System.Startup = DefaultSystem.Startup
	} else {
		cfg.System.Startup = startup
	}
}

func (cfg *Config) UpdateShutdown(shutdown string) {
	// Update shutdown message if provided
	if shutdown == "" {
		cfg.System.Shutdown = DefaultSystem.Shutdown
	} else {
		cfg.System.Shutdown = shutdown
	}
}
