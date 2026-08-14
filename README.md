[![CI](https://github.com/hendrik-ro/replGo/actions/workflows/ci.yml/badge.svg)](https://github.com/hendrik-ro/replGo/actions/workflows/ci.yml)

REPL Go
=======
A CLI REPL package for Go.

## About
A custom made REPL package for Go to be used as a library or standalone CLI tool. It provides a simple REPL loop with customizable prompt and input handling.

## Features
- Adding and removing custom commands
- Returning command list
- History tracking and listing

## Roadmap
- Default commands: exit, help

## Usage
To use the REPL package, import it into your Go project and create a new `Repl` instance with your desired prompt and input handler functions.

```go
import "github.com/hendrik-ro/replGo"

func main() {
	repl := replGo.NewRepl("")
	repl.Add(replGo.Command{
		Name: "echo",
		Description: "Echoes the input back to the user",
		Handler: func(args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	})
	repl.Run()
}
```

## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request on the [GitHub repository](https://github.com/hendrik-ro/replGo).

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
