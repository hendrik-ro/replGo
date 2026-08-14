package repl

import (
	"fmt"
)

func newREPL() *REPL {
	return &REPL{
		Config: Config{},
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
