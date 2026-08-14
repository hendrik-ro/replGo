package repl

func (r *REPL) Add(cmd Command) {
	// Adds a command to the REPL's command list
	r.Config.Commands = append(r.Config.Commands, cmd)
}

func (r *REPL) Remove(name string) {
	// Removes a command from the REPL's command list
	for i, cmd := range r.Config.Commands {
		if cmd.Name == name {
			r.Config.Commands = append(r.Config.Commands[:i], r.Config.Commands[i+1:]...)
			return
		}
	}
}

func (r *REPL) List() *[]Command {
	// Returns a pointer to the REPL's command list
	return &r.Config.Commands
}
