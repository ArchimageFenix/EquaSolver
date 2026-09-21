// File: main.go
// Purpose: entry point of the command-line program.
// Receives: arguments, standard input and standard output from the operating system.
// Previous stage: the operating system.
// Next stage: internal/cli (Run).
// Restrictions: no calculation, formatting or data entry; it only starts the CLI
//               and exits with its code.

package main

import (
	"os"

	"gauss/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdin, os.Stdout))
}
