// File: substitution_test.go
// Purpose: check back substitution and its failure flag.
// Receives: hand-made echelon matrices.
// Previous stage: substitution.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package substitution

import (
	"math"
	"testing"
)

func TestSolveBackSubstitution(t *testing.T) {
	echelon := [][]float64{{2, -1, 1, 3}, {0, 2.5, -1.5, 0.5}, {0, 0, 1.4, 4.2}}
	solution, ok := Solve(echelon)
	if !ok {
		t.Fatal("unexpected failure")
	}
	for i, want := range []float64{1, 2, 3} {
		if math.Abs(solution[i]-want) > 1e-9 {
			t.Errorf("x%d = %g, want %g", i+1, solution[i], want)
		}
	}
}

func TestSolveSingleEquation(t *testing.T) {
	solution, ok := Solve([][]float64{{2, 4}})
	if !ok || solution[0] != 2 {
		t.Errorf("got %v, %v", solution, ok)
	}
}

func TestSolveReportsNonFiniteResult(t *testing.T) {
	if _, ok := Solve([][]float64{{1e-300, 1e300}}); ok {
		t.Error("an overflow must be reported")
	}
}
