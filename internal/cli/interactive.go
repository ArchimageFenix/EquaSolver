// File: interactive.go
// Purpose: ask n and the equations one by one from the keyboard, with help and retry.
// Receives: a prompter (console input and output).
// Previous stage: menu.go (option 1).
// Next stage: run.go, which sends the model.System to the solver.
// Restrictions: does not calculate; only checks the text format of each answer;
//               never allocates memory based on n alone (rows are added as they arrive).

package cli

import (
	"gauss/internal/model"
	"gauss/internal/numeric"
)

func readSystemTyped(p prompter) (model.System, error) {
	n, err := askUntilValid(p, countQuestion, parseCount)
	if err != nil {
		return model.System{}, err
	}
	rows, err := askRows(p, n)
	if err != nil {
		return model.System{}, err
	}
	return model.FromAugmentedRows(rows), nil
}

func askRows(p prompter, n int) ([][]float64, error) {
	rows := [][]float64{}
	for i := 1; i <= n; i++ {
		row, err := askUntilValid(p, rowQuestion(i, n), func(text string) ([]float64, error) {
			return numeric.ParseRow(text, n+1)
		})
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}
