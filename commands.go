package repl

import (
	"fmt"
	"os"
)

var DefaultCommands = []Command{
	DefaultHelp,
	DefaultExit,
}

var DefaultHelp = Command{
	Name:        "help",
	Description: "Display this help message",
	Handler: func(cfg *Config, args []string) {
		fmt.Println("Available commands:")
		for _, cmd := range cfg.Commands {
			fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
		}
	},
}
var DefaultExit = Command{
	Name:        "exit",
	Description: "Exit the REPL",
	Handler: func(cfg *Config, args []string) {
		fmt.Println(cfg.System.Shutdown)
		os.Exit(0)
	},
}
