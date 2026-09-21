// File: solver_test.go
// Purpose: check the whole core on the three cases, edge cases and failures.
// Receives: hand-made systems.
// Previous stage: solver.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package solver

import (
	"math"
	"testing"

	"gauss/internal/model"
)

func systemOf(rows ...[]float64) model.System {
	return model.FromAugmentedRows(rows)
}

func TestSolveUniqueSolution(t *testing.T) {
	result, err := Solve(systemOf([]float64{1, 1, 1, 6}, []float64{2, -1, 1, 3}, []float64{1, 2, -1, 2}))
	if err != nil || result.Case != model.UniqueSolution {
		t.Fatalf("unexpected result: %v, %v", result.Case, err)
	}
	for i, want := range []float64{1, 2, 3} {
		if math.Abs(result.Solution[i]-want) > 1e-9 {
			t.Errorf("x%d = %g, want %g", i+1, result.Solution[i], want)
		}
	}
}

func TestSolveInfiniteSolutions(t *testing.T) {
	result, err := Solve(systemOf([]float64{1, 1, 1, 6}, []float64{2, -1, 1, 3}, []float64{3, 0, 2, 9}))
	if err != nil || result.Case != model.InfiniteSolutions {
		t.Fatalf("unexpected result: %v, %v", result.Case, err)
	}
	if result.RankCoefficients != 2 || result.RankAugmented != 2 || result.FreeVariables != 1 {
		t.Errorf("unexpected ranks: %+v", result)
	}
	if result.Solution != nil {
		t.Error("infinite solutions must not carry a solution")
	}
}

func TestSolveNoSolutionPointsToContradictoryRow(t *testing.T) {
	result, err := Solve(systemOf([]float64{1, 1, 1, 6}, []float64{2, -1, 1, 3}, []float64{3, 0, 2, 10}))
	if err != nil || result.Case != model.NoSolution {
		t.Fatalf("unexpected result: %v, %v", result.Case, err)
	}
	if result.RankAugmented != 3 || result.ContradictoryRow != 2 {
		t.Errorf("unexpected classification: %+v", result)
	}
}

func TestSolveSingleEquation(t *testing.T) {
	result, err := Solve(systemOf([]float64{2, 4}))
	if err != nil || result.Case != model.UniqueSolution || result.Solution[0] != 2 {
		t.Errorf("unexpected result: %+v, %v", result, err)
	}
}

func TestSolveAllZeroCoefficients(t *testing.T) {
	result, err := Solve(systemOf([]float64{0, 0, 0}, []float64{0, 0, 5}))
	if err != nil || result.Case != model.NoSolution {
		t.Errorf("unexpected result: %v, %v", result.Case, err)
	}
	result, err = Solve(systemOf([]float64{0, 0, 0}, []float64{0, 0, 0}))
	if err != nil || result.Case != model.InfiniteSolutions {
		t.Errorf("unexpected result: %v, %v", result.Case, err)
	}
}

func TestSolveRejectsInvalidSystem(t *testing.T) {
	if _, err := Solve(model.System{N: 0}); err == nil {
		t.Error("an invalid system must be rejected")
	}
}

func TestSolveReportsNumericFailure(t *testing.T) {
	result, err := Solve(systemOf([]float64{1e308, 1e308, 1e308}, []float64{-1e308, 1e308, 1e308}))
	if err != nil || result.Case != model.NumericFailure {
		t.Fatalf("unexpected result: %v, %v", result.Case, err)
	}
	if result.Echelon != nil {
		t.Error("a numeric failure must not expose a partial matrix")
	}
}
