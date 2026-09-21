// File: formcheck.go
// Purpose: validate the form text on the server before anything reaches the core.
// Receives: the raw text of the matrix field (one equation per line, or ';').
// Previous stage: handlers.go.
// Next stage: handlers.go, which sends the model.System to solver.Solve.
// Restrictions: rejects invalid text with a clear message and never lets it continue
//               to the solver; no calculations; the page's own check is only a
//               convenience and is not trusted.

package web

import (
	"errors"

	"gauss/internal/model"
	"gauss/internal/numeric"
)

// checkForm parses the matrix text into a system or explains what is wrong.
func checkForm(text string) (model.System, error) {
	lines := numeric.SplitRows(text)
	if len(lines) == 0 {
		return model.System{}, errors.New("ingrese al menos una ecuación")
	}
	rows, err := numeric.ParseRows(lines, len(lines))
	if err != nil {
		return model.System{}, err
	}
	return model.FromAugmentedRows(rows), nil
}
