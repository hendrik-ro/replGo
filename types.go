package repl

type REPL struct {
	Config *Config
}

type Config struct {
	// Configuration for the REPL system.

	Commands []Command
	History  *History
	Prompt   string
	System   System
}

type System struct {
	// Startup and shutdown messages for the REPL system.
	Startup  string
	Shutdown string
}

type Command struct {
	// Command metadata for the REPL system.
	Name        string
	Description string
	// Handler function for the command.
	Handler func(Config, []string)
}

type History struct {
	// Maximum number of history entries to keep.
	MaxSize int
	// List of history entries.
	Entries *[]HistoryEntry
}

type HistoryEntry struct {
	// Command and arguments for the history entry.
	Command Command
	Args    []string
}
