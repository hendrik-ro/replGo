package repl

import (
	"fmt"

	typesREPL "github.com/hendrik-ro/replGo/module/types"
)

func newREPL(cfg typesREPL.Config) typesREPL.REPL {
	return typesREPL.REPL{
		Config: cfg,
	}
}

func (r *typesREPL.REPL) Run() {
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
