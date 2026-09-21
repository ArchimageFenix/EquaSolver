// File: file.go
// Purpose: read a system from a text file (first line n, then n lines of n+1 numbers).
// Receives: a file path (from -f or asked through the prompter).
// Previous stage: menu.go (option 2) or run.go (-f argument).
// Next stage: run.go, which sends the model.System to the solver.
// Restrictions: does not calculate; rejects a malformed file before the core sees it;
//               the line count is checked against n before any data is used.

package cli

import (
	"errors"
	"fmt"
	"os"

	"gauss/internal/model"
	"gauss/internal/numeric"
)

// readSystemFile reads and parses the system stored in the file.
func readSystemFile(path string) (model.System, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.System{}, fmt.Errorf("no se pudo leer el archivo: %w", err)
	}
	return systemFromLines(numeric.SplitRows(string(data)))
}

// readSystemFromPrompt asks for the path until a readable, well-formed file is given.
func readSystemFromPrompt(p prompter) (model.System, error) {
	return askUntilValid(p, pathQuestion, readSystemFile)
}

func systemFromLines(lines []string) (model.System, error) {
	if len(lines) == 0 {
		return model.System{}, errors.New("el archivo está vacío")
	}
	n, err := parseCount(lines[0])
	if err != nil {
		return model.System{}, err
	}
	if len(lines)-1 != n {
		return model.System{}, fmt.Errorf("se esperaban %d ecuaciones y el archivo tiene %d", n, len(lines)-1)
	}
	return systemFromEquations(lines[1:])
}
