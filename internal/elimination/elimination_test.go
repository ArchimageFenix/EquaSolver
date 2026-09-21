// File: elimination_test.go
// Purpose: check pivoting, recorded steps, rank and failure detection.
// Receives: hand-made systems.
// Previous stage: elimination.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package elimination

import (
	"math"
	"testing"

	"gauss/internal/model"
)

func systemOf(rows ...[]float64) model.System {
	return model.FromAugmentedRows(rows)
}

func TestRunSwapsRowsAndReachesFullRank(t *testing.T) {
	system := systemOf([]float64{1, 1, 1, 6}, []float64{2, -1, 1, 3}, []float64{1, 2, -1, 2})
	outcome := Run(system, 1e-12)
	if outcome.Failed || outcome.Rank != 3 {
		t.Fatalf("unexpected outcome: failed=%v rank=%d", outcome.Failed, outcome.Rank)
	}
	first := outcome.Steps[0]
	if first.SwappedWith != 1 || first.Pivot != 2 || len(first.Operations) != 2 {
		t.Errorf("unexpected first step: %+v", first)
	}
	if got := outcome.Echelon[2][2]; math.Abs(got-1.4) > 1e-9 {
		t.Errorf("last pivot = %g, want 1.4", got)
	}
}

func TestRunKeepsOriginalAndInputUntouched(t *testing.T) {
	system := systemOf([]float64{1, 1, 1, 6}, []float64{2, -1, 1, 3}, []float64{1, 2, -1, 2})
	outcome := Run(system, 1e-12)
	if outcome.Original[0][0] != 1 || system.A[0][0] != 1 {
		t.Error("the original data must not change")
	}
}

func TestRunReportsMissingPivot(t *testing.T) {
	system := systemOf([]float64{1, 1, 1, 6}, []float64{2, -1, 1, 3}, []float64{3, 0, 2, 9})
	outcome := Run(system, 1e-12)
	if outcome.Rank != 2 {
		t.Fatalf("rank = %d, want 2", outcome.Rank)
	}
	last := outcome.Steps[len(outcome.Steps)-1]
	if !last.NoPivot {
		t.Errorf("the last column should have no pivot: %+v", last)
	}
}

func TestRunDetectsNonFiniteValues(t *testing.T) {
	system := systemOf([]float64{1e308, 1e308, 1e308}, []float64{-1e308, 1e308, 1e308})
	outcome := Run(system, 1e-12)
	if !outcome.Failed {
		t.Error("an overflow must be reported as a failure")
	}
}
