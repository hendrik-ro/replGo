[![CI](https://github.com/hendrik-ro/replGo/actions/workflows/ci.yml/badge.svg)](https://github.com/hendrik-ro/replGo/actions/workflows/ci.yml)

REPL Go
=======
A CLI REPL package for Go.

## About
A custom made REPL package for Go to be used as a library or standalone CLI tool. It provides a simple REPL loop with customizable prompt and input handling.

## Features
- Default commands: exit, help
- Adding and removing custom commands
- Returning command list
- History tracking and listing
- Prompt customization
- Customization of system startup and shutdown messages

## Roadmap
- tbd

## Usage
To use the REPL package, import it into your Go project and create a new `Repl` instance with your desired prompt and input handler functions.

```go
import "github.com/hendrik-ro/replGo"

func main() {
	// Initialize the REPL
	repl := replGo.NewRepl()
	
	// Add custom commands
	repl.Add(replGo.Command{
		Name: "echo",
		Description: "Echoes the input back to the user",
		Handler: func(args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	})

	// Run the REPL
	repl.Run()
}
```

## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request on the [GitHub repository](https://github.com/hendrik-ro/replGo).

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
