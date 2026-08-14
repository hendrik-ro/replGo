package repl

import (
	"fmt"
)

func newREPL(cfg Config) *REPL {
	return &REPL{
		Config: cfg,
	}
}

func (r *REPL) Run() {
	for {
		fmt.Println("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("failed to read input: ", err)
			return
		}

	}
}
