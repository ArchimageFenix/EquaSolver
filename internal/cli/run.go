// File: run.go
// Purpose: run the CLI flow: arguments, data source, solver, presentation.
// Receives: the arguments, the console input and the console output.
// Previous stage: cmd/cli/main.go.
// Next stage: solver.Solve (with the model.System) and present.go (with the result).
// Restrictions: no calculations and no formatting rules; it only chains the CLI
//               pieces and returns the process exit code.

package cli

import (
	"fmt"
	"io"

	"gauss/internal/model"
	"gauss/internal/solver"
)

const (
	exitOK    = 0
	exitInput = 1
	exitUsage = 2
)

// Run executes the CLI and returns the exit code.
func Run(args []string, in io.Reader, out io.Writer) int {
	parsed, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(out, err)
		return exitUsage
	}
	system, err := obtainSystem(parsed, newPrompter(in, out))
	if err != nil {
		fmt.Fprintln(out, "Error:", err)
		return exitInput
	}
	result, err := solver.Solve(system)
	if err != nil {
		fmt.Fprintln(out, "Error:", err)
		return exitInput
	}
	printResult(out, result)
	return exitOK
}

// obtainSystem picks the source of the data: file, matrix argument or the menu.
func obtainSystem(parsed arguments, p prompter) (model.System, error) {
	switch {
	case parsed.file != "":
		return readSystemFile(parsed.file)
	case parsed.matrix != "":
		return systemFromMatrixArgument(parsed.matrix)
	default:
		return readInteractively(p)
	}
}
