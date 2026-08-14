package repl

import (
	"fmt"
	"strings"
)

func newREPL() *REPL {
	// Initialize the REPL with a default config
	cfg := Config{
		History: &History{
			MaxSize: 100,
			Entries: &[]HistoryEntry{},
		},
	}
	return &REPL{
		Config: &cfg,
	}
}

func (r *REPL) Run() {
	// Start the REPL loop
	for {
		fmt.Println("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("failed to read input: ", err)
			return
		}
		inputCmd := strings.Split(input, " ")
		cmd := strings.TrimSpace(strings.ToLower(inputCmd[0]))
		args := inputCmd[1:]

		var lastCmd = Command{}

		for _, c := range r.Config.Commands {
			if c.Name == cmd {
				c.Handler(args)
				lastCmd = c
				break
			}
		}
		r.Config.History.Add(HistoryEntry{
			Command: lastCmd,
			Args:    args,
		})
	}
}
