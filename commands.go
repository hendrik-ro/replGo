package repl

func (r *REPL) Add(cmd Command) {
	r.Config.Commands = append(r.Config.Commands, cmd)
}
