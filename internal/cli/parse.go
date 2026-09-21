// File: parse.go
// Purpose: turn CLI text (counts, equation lines, the -m argument) into a model.System.
// Receives: text lines coming from the keyboard, a file or the -m argument.
// Previous stage: interactive.go, file.go and run.go.
// Next stage: the same callers, which send the model.System to the solver.
// Restrictions: text format only; the numeric rules live in package numeric and the
//               domain rules in validation; no I/O.

package cli

import (
	"errors"
	"strconv"

	"gauss/internal/model"
	"gauss/internal/numeric"
)

func parseCount(text string) (int, error) {
	n, err := strconv.Atoi(text)
	if err != nil || n < 1 {
		return 0, errors.New("n debe ser un entero mayor que 0")
	}
	return n, nil
}

// systemFromEquations builds the system from one line per equation.
func systemFromEquations(lines []string) (model.System, error) {
	rows, err := numeric.ParseRows(lines, len(lines))
	if err != nil {
		return model.System{}, err
	}
	return model.FromAugmentedRows(rows), nil
}

// systemFromMatrixArgument builds the system from the -m argument (rows split by ';').
func systemFromMatrixArgument(text string) (model.System, error) {
	lines := numeric.SplitRows(text)
	if len(lines) == 0 {
		return model.System{}, errors.New("la matriz está vacía")
	}
	return systemFromEquations(lines)
}
