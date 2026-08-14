package repl

import (
	"fmt"
	"strings"
)

func NewREPL() *REPL {
	// Initialize the REPL with a default config
	cfg := DefaultConfig()
	return &REPL{
		Config: cfg,
	}
}

func (r *REPL) Run() {
	// Start the REPL loop
	fmt.Println(r.Config.System.Startup)
	for {
		fmt.Print(r.Config.Prompt)
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
				c.Handler(r.Config, args)
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
