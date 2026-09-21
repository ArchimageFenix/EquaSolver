// File: system.go
// Purpose: define the neutral System (n equations, n unknowns) and build it from rows.
// Receives: rows of n coefficients followed by one independent term.
// Previous stage: cli and web adapters (after parsing their own input).
// Next stage: solver (which validates it before any calculation).
// Restrictions: only data and its construction; no validation, no I/O.
//               FromAugmentedRows expects rows of n+1 values (numeric.ParseRows guarantees it).

package model

// System is the neutral input of the solver: A x = B with n equations.
type System struct {
	N int
	A [][]float64
	B []float64
}

// FromAugmentedRows builds a System from rows of n coefficients plus one independent term.
func FromAugmentedRows(rows [][]float64) System {
	n := len(rows)
	system := System{N: n, A: make([][]float64, n), B: make([]float64, n)}
	for i, row := range rows {
		system.A[i] = row[:n]
		system.B[i] = row[n]
	}
	return system
}
