package repl

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(r.Config.Prompt)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("failed to read input:", err)
			}
			return
		}

		input := scanner.Text()
		inputCmd := strings.Split(input, " ")
		if len(inputCmd) == 0 {
			continue
		}

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
