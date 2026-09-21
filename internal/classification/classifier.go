// File: classifier.go
// Purpose: decide the case (unique, infinite, none), the augmented rank and the
//          contradictory row from the echelon matrix.
// Receives: the echelon augmented matrix, the coefficient rank and the tolerance.
// Previous stage: solver (with the output of elimination).
// Next stage: solver (gets the classification and builds the Result).
// Restrictions: does not compute the solution; does not modify the matrix; no I/O.

package classification

import (
	"math"

	"gauss/internal/model"
)

// Outcome is the classification of a system.
type Outcome struct {
	Case             model.Case
	RankAugmented    int
	FreeVariables    int
	ContradictoryRow int // 0-based, -1 when it does not apply
}

// Classify compares the rank of the coefficients with the rank of the augmented matrix.
func Classify(echelon [][]float64, rank int, tolerance float64) Outcome {
	if row := findContradiction(echelon, rank, tolerance); row >= 0 {
		return Outcome{Case: model.NoSolution, RankAugmented: rank + 1, ContradictoryRow: row}
	}
	n := len(echelon)
	if rank == n {
		return Outcome{Case: model.UniqueSolution, RankAugmented: rank, ContradictoryRow: -1}
	}
	return Outcome{Case: model.InfiniteSolutions, RankAugmented: rank, FreeVariables: n - rank, ContradictoryRow: -1}
}

// findContradiction returns the first row below the rank whose independent term is
// not zero (a row 0 ... 0 | c), or -1 when there is none.
func findContradiction(echelon [][]float64, rank int, tolerance float64) int {
	n := len(echelon)
	for i := rank; i < n; i++ {
		if math.Abs(echelon[i][n]) > tolerance {
			return i
		}
	}
	return -1
}
