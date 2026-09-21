// File: classifier_test.go
// Purpose: check the three classifications and the contradictory row.
// Receives: hand-made echelon matrices.
// Previous stage: classifier.go.
// Next stage: none (test result).
// Restrictions: tests only; no shared state.

package classification

import (
	"testing"

	"gauss/internal/model"
)

func TestClassifyUniqueSolution(t *testing.T) {
	echelon := [][]float64{{2, -1, 1, 3}, {0, 2.5, -1.5, 0.5}, {0, 0, 1.4, 4.2}}
	got := Classify(echelon, 3, 1e-12)
	if got.Case != model.UniqueSolution || got.FreeVariables != 0 || got.ContradictoryRow != -1 {
		t.Errorf("unexpected classification: %+v", got)
	}
}

func TestClassifyInfiniteSolutions(t *testing.T) {
	echelon := [][]float64{{3, 0, 2, 9}, {0, -1, -0.3333, -3}, {0, 0, 0, 0}}
	got := Classify(echelon, 2, 1e-12)
	if got.Case != model.InfiniteSolutions || got.FreeVariables != 1 || got.RankAugmented != 2 {
		t.Errorf("unexpected classification: %+v", got)
	}
}

func TestClassifyNoSolutionReportsContradictoryRow(t *testing.T) {
	echelon := [][]float64{{3, 0, 2, 10}, {0, -1, -0.3333, -3.6667}, {0, 0, 0, -1}}
	got := Classify(echelon, 2, 1e-12)
	if got.Case != model.NoSolution || got.RankAugmented != 3 || got.ContradictoryRow != 2 {
		t.Errorf("unexpected classification: %+v", got)
	}
}

func TestClassifyTreatsTinyTermAsZero(t *testing.T) {
	echelon := [][]float64{{1, 1, 1}, {0, 0, 1e-15}}
	got := Classify(echelon, 1, 1e-12)
	if got.Case != model.InfiniteSolutions {
		t.Errorf("a term below the tolerance must count as zero: %+v", got)
	}
}
