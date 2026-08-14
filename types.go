package repl

type REPL struct {
	Config *Config
}

type Config struct {
	Commands []Command
	History  *History
	Prompt   string
}

type Command struct {
	Name        string
	Description string
	Handler     func(Config, []string)
}

type History struct {
	MaxSize int
	Entries *[]HistoryEntry
}

type HistoryEntry struct {
	Command Command
	Args    []string
}
