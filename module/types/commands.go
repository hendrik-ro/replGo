package typesREPL

type Command struct {
	Name        string
	Description string
	Handler     func()
}
