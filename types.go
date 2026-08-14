package repl

type Config struct {
	Commands []Command
}

type REPL struct {
	Config Config
}

type Command struct {
	Name        string
	Description string
	Handler     func([]string)
}
