// File: substitution.go
// Purpose: back substitution on an upper triangular augmented matrix.
// Receives: the echelon augmented matrix of a system with a unique solution.
// Previous stage: solver (only when the classification is a unique solution).
// Next stage: solver (gets the solution vector or a failure flag).
// Restrictions: never called for infinite or missing solutions; does not modify the
//               matrix; no I/O; reports failure instead of returning non-finite values.

package substitution

import "gauss/internal/numeric"

// Solve returns the unknowns, or false if a non-finite value appears.
func Solve(echelon [][]float64) ([]float64, bool) {
	n := len(echelon)
	solution := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		value := (echelon[i][n] - knownTerms(echelon[i], solution, i+1, n)) / echelon[i][i]
		if !numeric.IsFinite(value) {
			return nil, false
		}
		solution[i] = value
	}
	return solution, true
}

// knownTerms adds row[j]*solution[j] for j in [from, to).
func knownTerms(row, solution []float64, from, to int) float64 {
	sum := 0.0
	for j := from; j < to; j++ {
		sum += row[j] * solution[j]
	}
	return sum
}
